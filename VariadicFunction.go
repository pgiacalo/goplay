package main

import "fmt"

/*
The variadic parameter must be the last (or only) parameter
in the input parameter list. You indicate it with three
dots (…) before the type. The variable that’s created within
the function is a slice of the specified type.
*/
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(5, 1, 2, 3))
}
