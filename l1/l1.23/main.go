package main

import "fmt"

func deleteElement[T any](sl []T, n int) []T{
	if n < 0 || n >= len(sl) {
		fmt.Printf("index %v out of range for slice %v\n", n, sl)
		return sl
	}

	copy(sl[n:], sl[n+1:])

	var zero T
	sl[len(sl)-1] = zero

	return sl[:len(sl)-1]
}

func main() {
	slInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slFloat32 := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	slString := []string{"first", "second", "third", "fourth", "fifth"}
	n := 2

	slInt = deleteElement(slInt, n)
	slFloat32 = deleteElement(slFloat32, n)
	slString = deleteElement(slString, n)

	fmt.Println(slInt)
	fmt.Println(slFloat32)
	fmt.Println(slString)
}
