package main

import (
	"flag"
	"fmt"
)

func main() {
	var number int64
	var bitPosition int
	var selector int

	flag.Int64Var(&number, "n", 5, "enter a number")
	flag.IntVar(&bitPosition, "b", 0, "enter a bit position")
	flag.IntVar(&selector, "s", 0, "select an operation: 0 - Set 0; 1 - Set 1; 2 - invert bit")
	flag.Parse()
	switch selector {
	case 0:
		fmt.Println(SetZero(number, bitPosition))
	case 1:
		fmt.Println(SetOne(number, bitPosition))
	case 2:
		fmt.Println(BitConvertor(number, bitPosition))
	default:
		fmt.Println("something wrong")
	}
}

func BitConvertor(number int64, i int) int64 {
	number ^= 1 << i
	return number
}

func SetOne(number int64, i int) int64 {
	number |= 1 << i
	return number
}

func SetZero(number int64, i int) int64 {
	number &^= 1 << i
	return number
}
