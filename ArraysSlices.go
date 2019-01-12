package main

import(
	"fmt"
)

type LinesOfText [][]byte

type Transform [3][3]float64  // A 3x3 array, really an array of arrays.

func init(){
	fmt.Println("ArraySlices init() called")
}

func main(){

	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Printf("%v\n", text)
	fmt.Printf("%s\n", text)

	var matrix Transform
	fmt.Printf("%v\n", matrix)

}