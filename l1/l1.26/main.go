package main

import (
	"fmt"
	"strings"

)

func AllUnique(s string) bool {
	m := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, i := range s {
		if m[i] {
			return false
		}
		m[i] = true
	}
	
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"

	fmt.Println(s1, AllUnique(s1))
	fmt.Println(s2, AllUnique(s2))
}