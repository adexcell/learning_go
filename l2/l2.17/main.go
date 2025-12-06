package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Config struct {
	host    string
	port    string
	timeout time.Duration
}

func (c *Config) getAddress() string {
	return net.JoinHostPort(c.host, c.port)
}

func NewConfig() *Config {
	host := flag.String("h", "google.com", "host")
	port := flag.String("p", "80", "port")
	timeout := flag.Duration("t", (time.Second * 10), "timeout")
	flag.Parse()

	return &Config{
		host:    *host,
		port:    *port,
		timeout: *timeout,
	}
}

func main() {
	cfg := NewConfig()
	done := make(chan struct{})

	conn, err := net.DialTimeout("tcp", cfg.getAddress(), cfg.timeout)
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Print... ~ ")
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println("Connection closed by server")
		done <- struct{}{}
	}()

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Println("Error:", err)
		}
		done <- struct{}{}
	}()

	<-done

}
