package main

import "fmt"

func DeleteElement[T any](sl []T, n int) {
	_ = copy(sl[n:], sl[n+1:])
	sl = sl[:len(sl)-1]
	fmt.Println(sl)
}

func main() {
	slInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slFloat := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	slString := []string{"first", "second", "third", "fourth", "fifth"}
	n := 2
	DeleteElement(slInt, n)
	DeleteElement(slFloat, n)
	DeleteElement(slString, n)
}
