//This code was created after watching Rob Pike's youtube video on GoRoutines.
//https://www.youtube.com/watch?v=f6kdp27TYZs
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// --------------------------------------------------------------------------
	//Mantra: Don't communicate by sharing memory, share memory by communicating.
	// --------------------------------------------------------------------------
	//	- go channels are used for this communication

	//Channels block on send and on receive until content arrives.
	//Therefore, channels communicate and synchronize in a single operation.
	//This is a fundamental concept to channels.
	//Note: Buffered channels do NOT synchronize (requires more subtle reasoning).
	//		-- Buffered channels are more like mailboxes

	//create a new channel for type string
	c := make(chan string)

	//go routines start functions that run independently (like the unix &).
	//go routines have their own stack that grows/shrinks as required.
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) //receive msg from channel
	}
	fmt.Println("You're boring: I'm leaving.")

	d := generator("boring!") //call function that returns a channel
	fmt.Println("Start using generator()")
	for i := 0; i < 5; i++ {
		fmt.Printf("Generate: %q\n", <-d) //receive msg from channel
	}
	fmt.Println("Done using generator()")

	//2 generators (used in a way that return dependent, sequential results)
	joe := generator("Joe!") //call function that returns a joe channel
	ann := generator("Ann!") //call function that returns a ann channel
	fmt.Println("Start using joe and ann generators() -- with sequential results")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe) //receive and print a msg from joe channel
		fmt.Println(<-ann) //then receive and prinit a msg from ann channel
	}
	fmt.Println("Done using both joe and ann generators()")

	//multiplexer to mix the output of 2 channels non-sequentially
	fmt.Println("Start using multiplexer for joe and ann generators()")
	e := multiplexer(generator("Joe"), generator("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-e) //receive msg from the multiplexer channel the output from both Joe and Ann channels
	}
	fmt.Println("Done using multiplexer for joe and ann generators()")

	//multiplexer (using select) to mix the output of 2 channels non-sequentially
	fmt.Println("Start using multiplexer with select for joe and ann generators()")
	f := multiplexerUsingSelect(generator("Joe"), generator("Ann"))
	for {
		select {
		case s := <-f:
			fmt.Println(s)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("You're too slow.")
			fmt.Println("Done using multiplexer with select for joe and ann generators()")
			return
		}
	}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-f) //receive msg from the multiplexer channel the output from both Joe and Ann channels
	// }
	// fmt.Println("Done using multiplexer with select for joe and ann generators()")

}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) //send msg into channel
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func generator(msg string) <-chan string {
	c := make(chan string)
	//start the generator function that puts data into the channel (to communicate with caller)
	go func() { //we launch the go routine from inside the generator function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) //send msg into channel
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}() //anonymous function literal

	return c //return the channel
}

//-- Non-sequential, independent go channels --
//A multiplexer takes multiple input channels and combines their independent outputs into a single/new output channel.
//This is also called a "fan"in" function. This allows multiple channel outputs to flow in non-sequential order.
//
//Channel 1 output \
//             	    \
//			    	 --> Returned Channel Output (combining/mixing Joe and Ann channel outputs)
//             		/
//Channel 2 output /
func multiplexer(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { //start collecting outputs from channel input1
		for {
			c <- <-c1
		}
	}() //anonymous function literal
	go func() { //start collecting outputs from channel input2
		for {
			c <- <-c2
		}
	}() //anonymous function literal
	return c //return the multiplex channel
}

//-- Non-sequential, independent go channels --
//A multiplexer takes multiple input channels and combines their independent outputs into a single/new output channel.
//This is also called a "fan"in" function. This allows multiple channel outputs to flow in non-sequential order.
//
//Channel 1 output \
//             	    \
//			    	 --> Returned Channel Output (combining/mixing Joe and Ann channel outputs)
//             		/
//Channel 2 output /
func multiplexerUsingSelect(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
				// //send a timeout message (if the other channels are slow)
				// case <-time.After(200 * time.Millisecond):
				// 	c <- fmt.("timeoutYou're too slow.")
			}
		}
	}()
	return c
}
