package main

import "fmt"

func binSearch(arr []int, n int) int {
	left := 0
	right := len(arr)-1
	mid := (left + right) / 2

	for left <= right{
		if n == arr[mid] {
			return mid
		} else if n < arr[mid] {
			right = mid - 1
			mid = (left + right) / 2
		} else {
			left = mid  + 1
			mid = (left + right) / 2
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binSearch(arr, 1))
	fmt.Println(binSearch(arr, 20))
}
