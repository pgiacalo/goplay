package main

import (
	"fmt"
	"time"
)

// Creates 2 channels and sends a message back and forth between them.
// The purpose is to discover the speed of channel communication.
// The result for this simple test is a message rate of ~6.5 million messages per second!
func main() {
	//create 3 channels (two for message passing and one for clock ticks)
	a := make(chan string, 10)      //channel for string messages (size 10)
	b := make(chan string, 10)      //channel for string messages (size 10)
	c := time.Tick(1 * time.Second) //channel for time ticks sent once per second

	count := 0 //message count between clock ticks
	total := 0 //total message count

	start := time.Now()
	stop := start.Add(5 * time.Second) //we'll stop counting after 5 seconds

	//put the message into channel a
	a <- "Now is the time for all good men to come to the aid of their country!"

LabelA:
	for {
		select {
		case msg := <-a:
			count++
			total++
			b <- msg
		case msg := <-b:
			count++
			total++
			a <- msg
		case tick := <-c:
			duration := tick.Sub(start)
			fmt.Printf("Count: %v, tick: %.1f\n", count, duration.Seconds())
			count = 0
			if tick.After(stop) {
				fmt.Printf("We're finished\n")
				break LabelA
			}
		}
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Total: %v, duration: %.1f, rate: %v\n", total, duration.Seconds(), float64(total)/duration.Seconds())
}
