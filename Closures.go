package main

//https://tour.golang.org/moretypes/25
//Function closures
//Go functions may be closures.
// DEFINITION: A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
// These referenced variables persist between calls to the closure.
//
//For example, the adder function returns a closure. Each closure is bound to its own sum variable

import "fmt"

//add() returns a function that takes and int and returns and int
//it is stateful, since the sum variable's value persists between calls (since this is a closure)
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//pos and neg are variables that refer to the 2 functions/closures returned by adder()
	//note how the values of sum persist (independently) for each of the 2 closures.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
