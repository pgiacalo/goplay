package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	numCPUs := runtime.NumCPU()
	println("CPUs =", numCPUs)

	runtime.GOMAXPROCS(numCPUs) //sets the max number of processors this will use

	// run 8 tasks in parallel
	go generatePrimes(1)
	go generatePrimes(2)
	go generatePrimes(3)
	go generatePrimes(4)
	go generatePrimes(5)
	go generatePrimes(6)
	go generatePrimes(7)
	go generatePrimes(8)

	//sleep to keep this main thread alive while the go routines are running
	sleepDur, _ := time.ParseDuration("2s")
	time.Sleep(sleepDur)
	fmt.Printf("Time main thread was alive: %v\n", sleepDur)
}

func generatePrimes(id int) []int {
	start := time.Now()
	const N = 20000000
	var x, y, n int
	nsqrt := math.Sqrt(N)

	is_prime := [N]bool{}

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			//fmt.Printf("id:%d, prime:%d\n", id, x)
			primes = append(primes, x)
		}
	}

	fmt.Printf("Duration #%v: %v\n", id, time.Since(start))

	return primes

	// primes is now a slice that contains all the
	// primes numbers up to N

	// let's print them
	//for _, x := range primes {
	//	fmt.Println(x)
	//}
}
