package main

import "fmt"


func main(){
	var x float64 = 123456.7890001
	fmt.Printf("%.2f\n", x)	//123456.79
	fmt.Printf("%8.2f\n", x)	//123456.79
	fmt.Printf("%8.2g\n", x)	//1.2e+05
	fmt.Printf("%e\n", x)	//1.234568e+05
	fmt.Printf("%E\n", x)	//1.234568E+05

	fmt.Printf("%.2f, %T\n", x, x) //123456.79, float64

	str := fmt.Sprintf("Sprintf = %.2f\n", x)
	fmt.Print(str)	//Sprintf = 123456.79

	var a [3][3]int
	fmt.Printf("Type of a: %T\n", a)	//Type of a: [3][3]int
	fmt.Printf("a = %v\n", a)			//a = [[0 0 0] [0 0 0] [0 0 0]]
}
