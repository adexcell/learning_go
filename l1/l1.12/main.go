package main

import (
	"fmt"
)

func MakeStringSet(a []string) []string {
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}

	result := make([]string, 0, len(a))

	for k := range m {
		result = append(result, k)
	}

	return result
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeStringSet(a))
}
