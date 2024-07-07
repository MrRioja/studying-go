package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func getArea(s shape) float64 {
	return s.area()
}

func getPerimeter(s shape) float64 {
	return s.perimeter()
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := rectangle{10, 15}
	fmt.Printf("The area is %0.2f and the perimeter is %0.2f\n", getArea(r), getPerimeter(r))
	c := circle{10}
	fmt.Printf("The area is %0.2f and the perimeter is %0.2f\n", getArea(c), getPerimeter(c))
}
