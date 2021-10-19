package main

import (
	"fmt"
	"time"
)

func statusUpdate() string {
	return "tick"
}

func main() {
	//func Tick(d Duration) <-chan Time
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
}
