package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	do()
}

func do() {
	defer printStack()

	fmt.Println("f() called")
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
