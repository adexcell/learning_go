package main

import (
	"bufio"
	"fmt"
	"os"
)


func reverse(runs []rune, st, fn int) {
	for i, j := st, fn; i < j; i, j = i+1, j-1 {
		runs[i], runs[j] = runs[j], runs[i]
	}
}

func LnFlipper(ln string) string {
	ln = ln[:len(ln)-1]
	r := []rune(ln)
	reverse(r, 0, len(r)-1)
	st := 0

	for i := 0; i < len(r); i++ {
		if r[i] == r[len(r)-1] || r[i] == rune(' ') {
			reverse(r, st, i)
			st = i + 1
			i++
		}
	}
	return string(r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter:")
	ln, _ := reader.ReadString('\n')
	fmt.Println(LnFlipper(ln))
}
