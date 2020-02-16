package main

import "math"

import "fmt"

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Printf("Area: %f\n", s.area())
}

func main() {
	c := circle{5}
	info(c)
}
