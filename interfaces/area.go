package main

import "fmt"

type shape interface {
	getArea() float64
}

type traingle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s traingle) getArea() float64 {
	return 0.5 * s.base * s.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{10}
	tr := traingle{10, 20}
	printArea(sq)
	printArea(tr)
}
