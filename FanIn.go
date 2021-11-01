/**
 * The Fan In pattern is used when there are MULTIPLE PRODUCERS and ONE CONSUMER.
 */
package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, startWith int, d time.Duration) {
	for {
		startWith++
		ch <- startWith
		time.Sleep(d)
	}
}

func consumer(ch chan int) {
	for {
		fmt.Printf("consumer: %d\n", <-ch)
	}
}

func main() {
	produced := make(chan int)
	consumed := make(chan int)

	//start multiple producers with each sending messages to the "produced" channel
	go producer(produced, 0, (100 * time.Millisecond))
	go producer(produced, 10000, (200 * time.Millisecond))
	go producer(produced, 10000000, (300 * time.Millisecond))
	//	close(produced)

	//start one consumer that receives messages from the "consumed" channel
	go consumer(consumed)

	//Create the "fan-in" by sending all messages from the
	//"produced" channel into the "consumed" channel
	for {
		consumed <- <-produced
	}

}
