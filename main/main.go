package main

import (
	"fmt"
)

var debug = true

func main() {
	if debug {
		fmt.Println("main starting...")
	}
	go Start()
}
