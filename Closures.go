package main

import (
	"fmt"
)

func main() {
	b := html_tag("<b>", "</b>")
	i := html_tag("<i>", "</i>")
	p := html_tag("<p>", "</p>")
	fmt.Println(b("bold"))      //<b>bold</b>
	fmt.Println(i("italic"))    //<i>italic</i>
	fmt.Println(p("paragraph")) //<p>paragraph</p>

	c1 := count(0)
	c2 := count(10)
	fmt.Println(c1(), c1()) //1 2
	fmt.Println(c1(), c1()) //3 4
	fmt.Println(c2(), c2()) //11 12

	sum := adder()
	fmt.Println(sum(5), sum(10)) //5 15

}

func html_tag(opentag string, closetag string) func(string) string {
	//definition: 	a closure is a function that remembers the variable(s)
	//				from the environment in which it was created.
	f := func(str string) string {
		return opentag + str + closetag
	}
	return f
}

func count(start int) func() int {
	//definition: 	a closure is a function that remembers the variable(s)
	//				from the environment in which it was created.
	f := func() int {
		start++
		return start
	}
	return f
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
