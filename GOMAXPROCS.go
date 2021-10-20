package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("ENV says GOMAXPROCS: :%s: \n", os.Getenv("GOMAXPROCS"))
	fmt.Printf("runtime says MAXPROCS = %d \n", runtime.NumCPU())
}
