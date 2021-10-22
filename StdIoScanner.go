// Scans lines of text entered from Stdin.
// Reports the count of repeated lines.
// Exits when a blank line is submitted.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	prompt := "Enter lines of data:"
	fmt.Println(prompt)
	//	for fmt.Println(prompt); input.Scan(); fmt.Println(prompt) {
	for {
		input.Scan()
		//get the line and trim whitespace
		line := strings.TrimSpace(input.Text())
		//exit when a blank line is entered
		if len(line) == 0 {
			break
		}
		counts[line]++
		//break out if the line is empty
	}
	fmt.Printf("count--%v--\n", counts)

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
