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

	fmt.Printf("title: %v, words:%v, isbn:%v\n", b.text.title, b.text.wordCount, b.isbn)
	//title: Romeo and Juliet, words:11657, isbn:1097834

	fmt.Printf("title: %v, words:%v, isbn:%v\n", b.title, b.wordCount, b.isbn)
	//title: Romeo and Juliet, words:11657, isbn:1097834

	fmt.Printf("%# v\n", pretty.Formatter(b))
	/* pretty.Formatter output
	main.book{
	    text: main.text{title:"Romeo and Juliet", wordCount:11657},
	    isbn: "1097834",
	}
	*/
}
