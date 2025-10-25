package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}

	pivotLeft := arr[left]
	pivotRight := arr[right]
	l := left + 1
	r := right - 1
	i := l

	for i <= r {
		if arr[i] < pivotLeft {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		} else if arr[i] > pivotRight {
			arr[i], arr[r] = arr[r], arr[i]
			r--
			i--
		}
		i++
	}

	l--
	r++
	arr[l], arr[left] = arr[left], arr[l]
	arr[r], arr[right] = arr[right], arr[r]

	quickSort(arr, left, l-1)
	quickSort(arr, l+1, r-1)
	quickSort(arr, r+1, right)

}

func slGen(length int, maxValue int) []int {
	sl := make([]int, 0, length)

	for range length {
		n := rand.IntN(maxValue)
		sl = append(sl, n)
	}

	return sl
}

func main() {
	var length, maxValue int
	flag.IntVar(&length, "l", 10, "enter slice length")
	flag.IntVar(&maxValue, "m", 100, "enter max value")
	flag.Parse()

	sl := slGen(length, maxValue)
	fmt.Println("not sorted:", sl)

	t0 := time.Now()

	quickSort(sl, 0, len(sl)-1)

	t1 := time.Now()

	fmt.Println("заняло времени", t1.Sub(t0))
	fmt.Println("sorted:", sl)
}
