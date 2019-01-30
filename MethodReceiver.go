package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

/*
Methods with pointer receivers CAN modify the value to which the receiver points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/
func (p *Person) MethodWithSideEffect(first string, last string) string {
	fmt.Printf("MethodWithSideEffect(): Type of Person, p %T\n", p)
	p.FirstName = first
	p.LastName = last
	return "middleName"
}

/*
Methods with value receivers can NOT modify the value to which the receiver points.
With a value receiver, the method operates on a copy of the original Vertex value.
(This is the same behavior as for any other value function argument)
*/
func (p Person) MethodWithoutSideEffect(first string, last string) Person {
	fmt.Printf("MethodWithoutSideEffect(): Type of Person, p %T\n", p)
	p.FirstName = first
	p.LastName = last
	return p
}

func (p Person) PrintFullName(msg string) {
	fmt.Printf("%s %s %s\n", msg, p.FirstName, p.LastName)
}

func changeParam(val *int, delta int) {
	fmt.Printf("1. val address=%p/n", val)
	*val = *val + delta
}

func noEffect(val int, delta int) {
	fmt.Printf("2. val address=%p/n", &val)
	val = val + delta
}

func main() {
	q := new(Person)
	fmt.Printf("Type of Person, q %T\n", q) //Type of Person, q *main.Person
	q.PrintFullName("1.")

	p := Person{
		"John",
		"Doe",
	}
	fmt.Printf("Type of Person, p %T\n", p) //Type of Person, p main.Person

	p.PrintFullName("2.")

	(&p).MethodWithSideEffect("Brian", "Greene")
	p.PrintFullName("3.") //p was changed

	m := p.MethodWithoutSideEffect("Little", "Richard")
	p.PrintFullName("4.") //p remains unchanged
	m.PrintFullName("5.") //a new person (m) was returned, while leaving p unchanged

	var x int
	fmt.Printf("0. val address=%p\n", &x)

	changeParam(&x, 6)
	fmt.Printf("x=%v\n", x)

	noEffect(x, 5)
	fmt.Printf("x=%v\n", x)
}
