package main

import (
	"fmt"
)

func main() {
	//m is a map of strings to float32s
	m := map[string]float32{
		"A": 1,
		"B": 2.56789,
		"C": 3.1415, //trailing comma is required
	}

	//the built-in delete() function removes items from maps (delete() does not return a value)
	delete(m, "C") //delete "C" from map m

	for key := range m {
		v, ok := m[key]
		if !ok {
			panic(fmt.Sprintf("No value found for key: %q", key))
		}
		fmt.Printf("key/value pair: %v = %v, ok = %v\n", key, v, ok)
	}

	//notice how the map returns a value even when the key doesn't exist.
	//it's necessary to always check the 'ok' value
	key := "nonexistent key"
	v, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("Key not found: %q, but a zero value is still returned: %v", key, v))
	}

}
