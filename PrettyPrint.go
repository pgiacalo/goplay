package main

import (
	"fmt"
	"time"

	"github.com/kr/pretty"
)

func main() {

	type Animal struct {
		Name   string
		Origin string
	}

	//A Bird has Animal characteristics
	type Bird struct {
		Animal //anonymous embedded struct - this simplifies dereferencing field values later
		Speed  int
		CanFly bool
	}

	//literal declaration to assign values
	b := Bird{
		Animal: Animal{Name: "emu", Origin: "Australia"}, //note how the embedded struct is exposed
		Speed:  48,
		CanFly: false,
	}

	fmt.Printf("name: %v, origin:%v, speed:%v, canfly:%v\n", b.Name, b.Origin, b.Speed, b.CanFly)
	//name: emu, origin:Australia, speed:48, canfly:false

	//Note: with anonymous embedding, the embedded field name is not needed when dereferencing (unless the outer struct also uses the same field name)
	fmt.Printf("name: %v, origin:%v, speed:%v, canfly:%v\n", b.Animal.Name, b.Animal.Origin, b.Speed, b.CanFly)
	//name: emu, origin:Australia, speed:48, canfly:false

	fmt.Printf("%# v\n", pretty.Formatter(b))

	fmt.Println(time.Now())
	/* pretty.Formatter output
	main.Bird{
	    Animal: main.Animal{Name:"emu", Origin:"Australia"},
	    Speed:  48,
	    CanFly: false,
	}
	*/
}
