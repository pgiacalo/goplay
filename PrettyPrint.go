package main

import (
	"fmt"
	"github.com/kr/pretty"
)

func main() {

	type Animal struct {
		Name   string
		Origin string
	}

	type Bird struct {
		Animal //embedded fields are NOT given a name (A bird has animal characteristics)
		Speed  int
		CanFly bool
	}

	//literal declaration to assign values
	b := Bird{
		Animal: Animal{Name: "emu", Origin: "Australia"}, //note how the embedded struct is exposed
		Speed:  48,
		CanFly: false,
	}

	fmt.Printf("%# v\n", pretty.Formatter(b))
}
