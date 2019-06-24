package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (sq square) getArea() float64 {
	return math.Pow(sq.sideLength, 2)
}

func (tg triangle) getArea() float64 {
	return 0.5 * tg.base * tg.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	s := square{2}
	t := triangle{2, 1}
	printArea(s)
	printArea(t)
}
