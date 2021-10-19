package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("Ticker started")

	done := make(chan bool)
	running := true

	go func() {
		for {
			select {
			case t := <-ticker.C:
				doPeriodicWork(t)
			case <-done:
				fmt.Println("Done processing")
				running = false
				return
			}
		}
	}()

	go exitAfter(5, done)

	//keep main running until we're done processing
	for running {
		time.Sleep(100 * time.Millisecond)
	}

	//release resources
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func exitAfter(seconds int, done chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}

func doPeriodicWork(t time.Time) {
	fmt.Println("Doing periodic work at", t)
	//simulate a process that takes 100ms
	milliseconds := 100
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
