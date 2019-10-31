package main

import (
	"fmt"
	"math"
)

// Shape : stores some operations to be used with different kinds of shapes
type Shape interface {
	area() float64
}

// Circle : generic proportional circle
type Circle struct {
	radius float64
}

// Rectangle : generic proportional rectangle
type Rectangle struct {
	width, height float64
}

// Square : generic square
type Square struct {
	width float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (s Square) area() float64 {
	return s.width * 4
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle1 := Circle{8}
	rect1 := Rectangle{8, 4}
	sqr1 := Square{2}

	fmt.Printf("Circle area: %f\n", getArea(circle1))
	fmt.Printf("Rectangle area: %f\n", getArea(rect1))
	fmt.Printf("Square area: %f\n", getArea(sqr1))
}
