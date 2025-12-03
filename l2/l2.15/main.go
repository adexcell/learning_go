package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	
	go func() {
		for {
			<- sigChan
			fmt.Println()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	u, _ := user.Current()

	for {
		showPrompt(u)
		if !(scanner.Scan()) {
			break
		}

		input := scanner.Text()

		if strings.TrimSpace(input) == "" {
			continue
		}

		if strings.Contains(input, "|") {
			handlePipeLine(input)
			continue
		} else {
			args := strings.Fields(input)
			command := args[0]

			handleCommand(u, command, args[1:])
		}
	}
}

func showPrompt(u *user.User) {
	fmt.Printf("%s:%s$ ", u.Username, showPWD())
}

func showPWD() string{
	wd, _ := os.Getwd()
	return wd
}

func killProccess(args []string) {
	if len(args) == 0 {
		return
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("PID должен быть числом")
		return
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	proc.Signal(os.Kill)
}

func handleCommand(u *user.User, cmdName string, args []string) {
	switch cmdName {
	case "echo":
		fmt.Println(strings.Join(args, " "))
	case "exit":
		os.Exit(1)
	case "cd":
		if len(args) == 0 {
			err := os.Chdir(u.HomeDir)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		
	case "pwd":
		fmt.Println(showPWD())
	case "kill":
		killProccess(args)
	default: 
		cmd := exec.Command(cmdName, args...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Ошибка выполнения: %v\n", err)
		}
	}
}

func handlePipeLine(input string) {
	parts := strings.Split(input, "|")

	var cmds []*exec.Cmd

	for _, part := range parts {
		args := strings.Fields(strings.TrimSpace(part))
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)
		cmds = append(cmds, cmd)
	}

	for i := 0; i < len(cmds)-1; i++ {
		currentCmd := cmds[i]
		nextCmd := cmds[i+1]

		stdoutPipe, err := currentCmd.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		nextCmd.Stdin = stdoutPipe
	}

	cmds[0].Stdin = os.Stdin
	cmds[len(cmds)-1].Stdout = os.Stdout

	for _, cmd := range cmds {
		cmd.Stderr = os.Stderr
	}

	for _, cmd := range cmds {
		err := cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка старта: %v\n", err)
			return
		}
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}
}