package main

import "fmt"

func main() {
	t := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(t)
	m := make(map[int][]float32)
	for _, i := range t {
		key := int(i) - int(i)%10
		m[key] = append(m[key], i)
	}
	fmt.Printf("%v", m)
}
