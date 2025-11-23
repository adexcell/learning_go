package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)


// Config is struct contains flags and filename
type Config struct {
	fields        string
	delimiter     string
	withSeparator bool
	filename      string
}

func parseArgs() Config {
	cfg := Config{}

	flag.StringVar(&cfg.fields, "f", "", "fields")
	flag.StringVar(&cfg.delimiter, "d", "\t", "delimeter")
	flag.BoolVar(&cfg.withSeparator, "s", false, "separator")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cfg.filename = args[len(args)-1]
	}

	return cfg
}

func parseFields(fields string) (map[int]bool, error) {
	m := make(map[int]bool, len(fields))
	columnNumbers := strings.Split(fields, ",")
	for _, i := range columnNumbers {
		if strings.Contains(i, "-") {
			rangeColumns := strings.Split(i, "-")
			f, s := rangeColumns[0], rangeColumns[1]
			start, err1 := strconv.Atoi(f)
			end, err2 := strconv.Atoi(s)

			if err1 != nil {
				return nil, err1
			}
			if err2 != nil {
				return nil, err2
			}

			for num := start; num <= end; num++ {
				m[num-1] = true
			}
			continue
		}
		num, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		m[num-1] = true
	}
	return m, nil
}

func cut(cfg Config, m map[int]bool) error {
	var scanner *bufio.Scanner
	if cfg.filename != "" {
		file, err := os.Open(cfg.filename)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, cfg.delimiter) {
			if cfg.withSeparator {
				continue
			}
			fmt.Println(line)
			continue
		}
		result := []string{}
		parts := strings.Split(line, cfg.delimiter)
		for i, v := range parts {
			if m[i] {
				result = append(result, v)
			}
		}
		resultString := strings.Join(result, cfg.delimiter)
		fmt.Println(resultString)
	}
	
	return scanner.Err()
}

func main() {
	cfg := parseArgs()

	m, err := parseFields(cfg.fields)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error:", err)
		os.Exit(1)
		return
	}

	err = cut(cfg, m)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error:", err)
		os.Exit(1)
		return
	}
}
