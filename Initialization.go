package main

import (
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// function to return the string representation of a Circle
func (c Circle) String() string {
	var s string = fmt.Sprintf("x:%v, y:%v, r:%v", c.x, c.y, c.r)
	return s
}

var f Circle
var g *Circle = new(Circle)
var h = Circle{x: 13, y: 14, r: 15}

func main() {
	//There are 3 different ways to create an object

	// 1) ---------
	c := Circle{x: 2, y: 3, r: 5}
	fmt.Printf("Circle c: %s\n", c.String()) //Circle c: x:2, y:3, r:5

	// 2) ---------
	var d Circle
	fmt.Printf("Circle d: %s\n", d.String()) //Circle d: x:0, y:0, r:0
	d.x = 10
	d.y = 11
	d.r = 12
	fmt.Printf("Circle d: %s\n", d.String()) //Circle d: x:10, y:11, r:12

	// 3) ---------
	e := new(Circle)
	fmt.Printf("Circle e: %s\n", e.String()) //Circle e: x:0, y:0, r:0
	e.x, e.y, e.r = 100, 200, 300
	fmt.Printf("Circle e: %s\n", e.String()) //Circle e: x:100, y:200, r:300

	// now access circle objects declared/instantiated outside of main()
	fmt.Printf("Circle f: %s\n", f.String()) //Circle f: x:0, y:0, r:0
	f.x, f.y, f.r = -1, -2, -3
	fmt.Printf("Circle f: %s\n", f.String()) //ircle f: x:-1, y:-2, r:-3

	fmt.Printf("Circle g: %s\n", g.String()) //Circle g: x:0, y:0, r:0
	g.x, g.y, g.r = -10, -20, -30
	fmt.Printf("Circle g: %s\n", g.String()) //Circle g: x:-10, y:-20, r:-30

	fmt.Printf("Circle h: %s\n", h.String()) //Circle h: x:13, y:14, r:15
	h.x, h.y, h.r = -100, -200, -300
	fmt.Printf("Circle h: %s\n", h.String()) //Circle h: x:-100, y:-200, r:-300

}
