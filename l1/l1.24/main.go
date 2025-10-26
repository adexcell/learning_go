package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance(otherPoint Point) float64 {
	return math.Pow(math.Pow(otherPoint.x-p.x, 2)+math.Pow(otherPoint.y-p.y, 2), 0.5)
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	p1 := NewPoint(1.0, 1.0)
	p2 := NewPoint(5.0, 5.0)
	distance := p1.Distance(p2)
	fmt.Println(distance)
}
