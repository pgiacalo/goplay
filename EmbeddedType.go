package main

import (
	"fmt"
)

type Person struct {
	Name string
}
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person		//an embedded type without a name represents the "is-a" relationship (an Android is-a type of Person)
	Model string
}

func main() {
	a := Android{Model: "cyborg", Person.Name: "3CPO"}
	//a.Name = "R2D2"
	a.Talk()

}