package main

import (
	"fmt"
)

type LinesOfText [][]byte // 2-dimensional slices

type Transform [3][3]float64 // 3x3 array, really an array of arrays.

func init() {
	fmt.Println("ArraySlices init() called")
}

func main() {

	//There are 4 ways to declare a slice:

	//1) Declare a slice and initialize with values
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("Type of s1=%T\n", s1) // Type of s1=[]int

	//2) NOT RECOMMENDED. initialize without values (memory is allocated)
	s2 := []int{}
	fmt.Printf("Type of s2=%T\n", s2) // Type of s2=[]int

	//3) Declare a slice but don't allocate memory just yet
	var s3 []int
	fmt.Printf("Type of s3=%T\n", s3) // Type of s3=[]int

	//4) Declare a slice and initialize without values with a length and (optional) capacity
	length := 5
	capacity := 10
	s4 := make([]int, length, capacity)
	fmt.Printf("Type of s4=%T\n", s4)         // Type of s4=[]int
	fmt.Printf("s4 length = %v\n", len(s4))   // s4 length = 5
	fmt.Printf("s4 capacity = %v\n", cap(s4)) // s4 capacity = 10

	s5 := make([]int, length)
	fmt.Printf("Type of s5=%T\n", s5)         // Type of s5=[]int
	fmt.Printf("s5 length = %v\n", len(s5))   // s5 length = 5
	fmt.Printf("s5 capacity = %v\n", cap(s5)) // s5 capacity = 5

	//an array is declared by specifying the length within the square brackets
	array := [3]string{"gold", "silver", "bronze"}
	fmt.Printf("Type of array=%T\n", array) // Type of array=[3]string

	var array2 = [3]string{}
	array2[0] = "gold"
	array2[1] = "silver"
	array2[0] = "bronze"
	fmt.Printf("Type of array2=%T\n", array2) // Type of array2=[3]string

	//a slice is declared without a size within the square brackets
	slice := []string{"gold", "silver", "bronze"}
	fmt.Printf("Type of slice=%T\n", slice) // Type of slice=[]string

	//Because slices are variable-length, it is possible to have each inner slice be a different length.

	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Printf("Type of text=%T\n", text) // Type of text=main.LinesOfText
	fmt.Printf("text=%v\n", text)         // text=[[78 111 119 32 105 115 32 116 104 ...
	fmt.Printf("text=%s\n", text)         // text=[Now is the time for all good gophers to bring some fun to the party.]

	var matrix Transform
	fmt.Printf("matrix=%v\n", matrix) // matrix=[[0 0 0] [0 0 0] [0 0 0]]

}
