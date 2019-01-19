package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

//implement the geometry interface for rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

//implement the geometry interface for circle
func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// implement a generalized function to handle any
// object that implements the geometry interface
func measure(g geometry) {
	fmt.Printf("measuring type %T: %+v\n", g, g)
	fmt.Printf("area = %v\n", g.area())
	fmt.Printf("perimeter = %v\n", g.perim())
}

func main() {
	r := rectangle{width: 3, height: 4}
	measure(r)

	c := circle{radius: 10}
	measure(c)
}
