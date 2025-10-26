package main

import "fmt"

// клиентский код
type Client struct {
}

func (c *Client) SomeUsefullFunc(p ClientInterface, msg string) {
	p.Print(msg)
}

// интерфейс клиента
type ClientInterface interface {
	Print(msg string)
}

// adapter
type Adapter struct {
	service *Service
}

func (a *Adapter) Print(msg string) {
	a.service.PrintSomeText(msg)
}

// сервис
type Service struct {
}

func (s *Service) PrintSomeText(msg string) {
	fmt.Println(msg)
}

func main() {

	msg := "hello world!"

	client := &Client{}
	service := &Service{}
	adapter := &Adapter{service: service}

	client.SomeUsefullFunc(adapter, msg)

}