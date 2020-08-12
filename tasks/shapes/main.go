package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triagle struct {
	height float64
	base   float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triagle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func main() {
	t := triagle{height: 2, base: 1}
	printArea(t)

	s := square{sideLenght: 2}
	printArea(s)
}
