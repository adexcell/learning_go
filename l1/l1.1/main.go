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
	Dream string
}

func (h *Human) Say(word string) {
	fmt.Println(word)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Abakar",
			Surname: "Aliev",
			Age: 31},
			Hobbie: "programming",
			Dream: "wb_ingener",
	}

	a.Say("hello world")

	fmt.Printf("Имя: %s, увлечение: %s\n", a.Name, a.Hobbie,)
}
