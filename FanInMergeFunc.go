/**
 * The Fan In pattern is used when there are MULTIPLE PRODUCERS and ONE CONSUMER.
 */
package main

import (
	"fmt"
	"time"
)

func Merge(out chan<- int, inA, inB <-chan int) {
	fmt.Println("Merge() called")
	for inA != nil || inB != nil {
		select {
		case v, ok := <-inA:
			if !ok {
				inA = nil
				fmt.Println("inA closed")
				continue
			}
			fmt.Println("inA open")
			out <- v
		case v, ok := <-inB:
			if !ok {
				inB = nil
				fmt.Println("inB closed")
				continue
			}
			fmt.Println("inB open")
			out <- v
			fmt.Println("out closed")

		}
	}
	fmt.Println("out closed")
	close(out)
}

func producer(ch chan<- int, startWith int, d time.Duration) {
	for {
		startWith++
		ch <- startWith
		time.Sleep(d)
	}
}

func consumer(ch <-chan int) {
	for {
		fmt.Printf("consumer: %d\n", <-ch)
	}
}

func main() {
	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)
	Merge(out, in1, in2)

	//start multiple producers with each sending messages to the "produced" channel
	go producer(in1, 0, (100 * time.Millisecond))
	go producer(in2, 10000, (200 * time.Millisecond))
	//	close(produced)

	//start one consumer that receives messages from the "consumed" channel
	go consumer(out)

}
