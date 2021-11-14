package main

import (
	"fmt"
)

func main() {
	s := "0123"
	s = reverse(s)
	fmt.Println(s)
	fmt.Println(reverse(s))
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
