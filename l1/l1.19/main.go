package main

import (
	"fmt"
	"unicode/utf8"
)

func Flipper(s string) string {
	length := utf8.RuneCountInString(s)
	newString := make([]rune, 0, length)
	r := []rune(s)
	for i := (length - 1); i >= 0; i-- {
		newString = append(newString, r[i])
	}
	return string(newString)
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(Flipper(s))
}
