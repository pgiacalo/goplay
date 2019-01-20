package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	userName := os.Getenv("USER")
	fmt.Println(userName)

}
