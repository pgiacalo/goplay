package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner(prompt string) {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" || input == "quit" || input == "q" {
			return
		}
		fmt.Printf("%v (type: %T)\n", scanner.Text(), scanner.Text())
	}
}

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Printf("%+v\n", argsWithProg)
	fmt.Printf("%+v\n", argsWithoutProg)

	arg := os.Args[0]
	fmt.Println(arg)

	scanner("Enter your input (type quit to exit):")
}
