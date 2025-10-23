package main

import "fmt"

func FindIntersection(a, b []int) []int {
	res := []int{}

	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 1, 5, 3}

	fmt.Println(FindIntersection(a, b))

}
