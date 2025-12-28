package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{base: 40, height: 30}
	sq := square{side: 40}
	printArea(tri)
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println("Area : ", s.getArea())
}
