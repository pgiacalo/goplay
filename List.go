package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	// New list.
	values := list.New()
	// Add 3 elements to the list.
	values.PushFront("bird")
	values.PushFront("cat")
	values.PushFront("snake")
	// Add 10 elements at the front.
	for i := 0; i < 10; i++ {
		// Convert ints to strings.
		values.PushFront(strconv.Itoa(i))
	}

	// Loop over container list.
	for temp := values.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}

	// Remove the oldest 5 elements
	for i := 0; i < 5; i++ {
		e := values.Back()
		values.Remove(e)
	}

}
