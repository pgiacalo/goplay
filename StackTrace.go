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

// In the code below, runtime.Stack() formats a stack trace of the
// calling goroutine into buf and returns the number of bytes written to buf.
// If all is true, Stack formats stack traces of all other goroutines into buf
// after the trace for the current goroutine.
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false) //if all = true, then all go routines are includes
	os.Stdout.Write(buf[:n])
}
