package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Person struct {
	name string
	age  int
}

//Reset() must be given a pointer in order to change/reset the Person (i.e., the receiver)
func (p *Person) Reset() {
	(*p).name = ""
	(*p).age = 0
}

var pool = sync.Pool{
	//Optional code to create an object if pool.Get() is called when the pool is empty
	New: func() interface{} {
		fmt.Printf("Pool is empty. Creating a new Person.\n")
		p := Person{}
		return p
	},
}

func main() {
	var persons [4]Person

	fmt.Printf("--- Create some objects and put them in the pool ---\n")

	// Create some objects and put them into the pool
	for i := 0; i < 3; i++ {
		p := Person{name: "Joe" + strconv.Itoa(i), age: i}
		pool.Put(p)
	}

	fmt.Printf("--- Get objects from the pool ---\n")

	// Get objects from the pool. When getting from a Pool, you need to cast
	for i := 0; i < 4; i++ {
		var p = pool.Get().(Person)
		fmt.Printf("%d) name=%q, age=%d\n", i, p.name, p.age)
		persons[i] = p
	}

	fmt.Printf("--- Modify the objects and put them back in the pool ---\n")

	// Modify them and put them back in the pool
	for i := 0; i < 4; i++ {
		p := persons[i]
		p.age = p.age + 1
		pool.Put(p)
	}

	fmt.Printf("--- Get objects from the pool ---\n")

	// Get objects from the pool. When getting from a Pool, you need to cast
	for i := 0; i < 4; i++ {
		var p = pool.Get().(Person)
		fmt.Printf("%d) name=%q, age=%d\n", i, p.name, p.age)
	}

	fmt.Printf("--- Reset the objects and put them back in the pool ---\n")

	// Reset them and put them back in the pool
	for i := 0; i < 4; i++ {
		p := persons[i]
		p.Reset()
		pool.Put(p)
	}

	fmt.Printf("--- Get objects from the pool ---\n")

	// Get objects from the pool. When getting from a Pool, you need to cast
	for i := 0; i < 4; i++ {
		var p = pool.Get().(Person)
		fmt.Printf("%d) name=%q, age=%d\n", i, p.name, p.age)
		persons[i] = p
	}

}
