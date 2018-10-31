package main

import (
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func area(s Shape) float64 {
	return s.area()
}

func main() {
	rect := Rectangle{width: 10, height: 15}
	circ := Circle{radius: 10}

	println(rect.area(), " --> Rectangle area")
	println(area(rect), "--> Rectangle Shape area")

	println(circ.area(), " --> Circle area")
	println(area(circ), " --> Circle Shape area")
}
