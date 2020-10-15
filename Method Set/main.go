package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

// Perimeter as pointer reciever only accepts *T
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Area as non pointer reciever works for T and *T
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	perimeter() float64
	area() float64
}

func info(s shape) {
	fmt.Println(s.area(), s.perimeter())
}

func main() {
	c := circle{1.0}
	info(&c) // Works
	//info(c)  //errors out - circle does not implement shape (perimeter method has pointer receiver)
}
