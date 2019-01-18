package main

import "fmt"

type Circle struct {
	x float64
	y float64
	r float64
}

func main() {
	var c Circle                     // c is of type main.Circle
	var d *Circle = new(Circle)      // d is a pointer of type *main.Circle
	var e = Circle{x: 0, y: 0, r: 5} // e is of type main.Circle

	fmt.Printf("Type of c = %T\n", c)
	fmt.Printf("Type of d = %T\n", d)
	fmt.Printf("Type of e = %T\n", e)

	c = e //ok
	c = e //ok
	//c = d 	//error - cannot use d (type *Circle) as type Circle in assignment
}
