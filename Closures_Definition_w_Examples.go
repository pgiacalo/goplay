/*
Closures
Functions declared inside of functions are special; they are closures.
This is a computer science word that means that functions declared
inside of functions are able to access and modify variables declared
in the outer function.
*/

//DEFINITION: A closure is an anonymous function that has access to the variable(s) from the environment in which it was CREATED.
//Function literals and closures. Function literals are closures: they inherit the scope of the function in which they are declared.

package main

import (
	"fmt"
)

func main() {
	b := html_tagger("<b>", "</b>")
	i := html_tagger("<i>", "</i>")
	p := html_tagger("<p>", "</p>")
	fmt.Println(b("bold"))      //<b>bold</b>
	fmt.Println(i("italic"))    //<i>italic</i>
	fmt.Println(p("paragraph")) //<p>paragraph</p>

	c1 := counter(1)
	c2 := counter(10)
	fmt.Println(c1(), c1(), c1(), c1()) //1 2 3 4
	fmt.Println(c2(), c2())             //10 11

	sum := adder()
	fmt.Println(sum(5), sum(10)) //5 15

}

func html_tagger(opentag string, closetag string) func(string) string {
	//closure
	f := func(str string) string {
		return opentag + str + closetag
	}
	return f
}

func counter(start int) func() int {
	//closure
	count := start - 1
	f := func() int {
		count++
		return count
	}
	return f
}

func adder() func(int) int {
	sum := 0
	//closure
	return func(x int) int {
		sum += x
		return sum
	}
}
