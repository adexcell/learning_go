package main

import "fmt"

func main() {
	num := 1
	word := "word"
	flag := true
	ch := make(chan int)
	PrintType(num)
	PrintType(word)
	PrintType(flag)
	PrintType(ch)
}

func PrintType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	}
}
