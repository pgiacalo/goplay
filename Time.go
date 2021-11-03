package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//Add 1 hours
	newT := t.Add(time.Hour * 1)
	fmt.Printf("Adding 1 hour\n: %s\n", newT)

	//Add 15 min
	newT = t.Add(time.Minute * 15)
	fmt.Printf("Adding 15 minute\n: %s\n", newT)

	//Add 10 sec
	newT = t.Add(time.Second * 10)
	fmt.Printf("Adding 10 sec\n: %s\n", newT)

	//Add 100 millisecond
	newT = t.Add(time.Millisecond * 10)
	fmt.Printf("Adding 100 millisecond\n: %s\n", newT)

	//Add 1000 microsecond
	newT = t.Add(time.Millisecond * 10)
	fmt.Printf("Adding 1000 microsecond\n: %s\n", newT)

	//Add 10000 nanosecond
	newT = t.Add(time.Nanosecond * 10000)
	fmt.Printf("Adding 1000 nanosecond\n: %s\n", newT)

	//Add 1 year 2 month 4 day
	newT = t.AddDate(1, 2, 4)
	fmt.Printf("Adding 1 year 2 month 4 day\n: %s\n", newT)
}
