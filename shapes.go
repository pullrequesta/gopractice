package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (s square) perimeter() float64 {
	return 2*s.length + 2*s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s shape) {
	fmt.Printf("area of %T is: %0.2f\n", s, s.area())
	fmt.Printf("perimeter of %T is: %0.2f\n", s, s.perimeter())
}
