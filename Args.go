package main

import (
	"fmt"
	"os"
	"strings"
)

func main1() {
	var s, sep string
	//skip the first Arg since it is the name of the function (i.e., main)
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	var c rune = 240
	fmt.Printf("%c\n", c)
}