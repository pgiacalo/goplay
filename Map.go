package main

import (
	"fmt"
)

func main() {
	m := map[string]float32{
		"A": 1.000001,
		"B": 2.56789,
	}
	for key := range m {
		v, ok := m[key]
		fmt.Printf("key/value pair: %v = %v, ok = %v\n", key, v, ok)
	}

	//notice how the map returns a value even when the key doesn't exist.
	//it's necessary to always check the 'ok' value
	key := "nonexistent key"
	v, ok := m[key]
	if !ok {
		panic("No value found for key")
	}
	fmt.Printf("key/value pair: %v = %.2g, ok = %t\n", key, v, ok)

}
