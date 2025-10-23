package main

import "fmt"

func main() {
	a := 4
	b := 10

	fmt.Printf("До обмена: а = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После обмена: а = %d, b = %d\n", a, b)

	c := 4
	d := 10

	fmt.Printf("До обмена: c = %d, d = %d\n", c, d)

	c = c - d
	d = d + c
	c = d - c

	fmt.Printf("После обмена: c = %d, d = %d\n", c, d)

	e := 2
	f := 3

	fmt.Printf("До обмена: e = %d, f = %d\n", e, f)

	e, f = f, e

	fmt.Printf("После обмена: e = %d, f = %d\n", e, f)

}
