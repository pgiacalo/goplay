package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

// Example demonstrates the use of the trace package to trace
// the execution of a Go program. The trace output will be
// written to the file trace.out that can be analyzed afterwards
// by the "go tool trace" command (and a Chrome browser).
// see: https://blog.gopheracademy.com/advent-2017/go-execution-tracer/
func main() {
	//setup a trace file to receive the trace output
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	//ok, now that we have a trace file, start the trace!
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// your program here
	DoThis()
	DoThat()
	DoMore()
}

func DoThis() {
	fmt.Printf("DoThis() this function will be traced")
}

func DoThat() {
	fmt.Printf("DoThat() this function will be traced")
}

func DoMore() {
	fmt.Printf("DoMore() this function will be traced")
}
