package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

type Action struct {
	// embedded struct
	Human
	Hobbie string
	Dream  string
}

func (h *Human) Say(word string) {
	fmt.Println(word)
}

func main() {
	a := Action{
		Human: Human{
			Name:    "Abakar",
			Surname: "Aliev",
			Age:     31},
		Hobbie: "programming",
		Dream:  "become a go senior developer",
	}

	a.Say("hello world")

	fmt.Printf("Имя: %s, увлечение: %s\n, цель: %s", a.Name, a.Hobbie, a.Dream)
}
