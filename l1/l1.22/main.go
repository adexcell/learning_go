package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	a = *big.NewInt(1_000_000_000)
	b = *big.NewInt(1_000_000_000_000)
	add := new(big.Int).Add(&a, &b)
	sub := new(big.Int).Sub(&a, &b)
	mul := new(big.Int).Mul(&a, &b)
	div := new(big.Int).Div(&a, &b)
	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
