/**
 * The Fan Out pattern is used when there are MULTIPLE WORKERS and a given set of work to be performed.
 * This code example applies ONLY if the jobs to be performed can be done in any order.
 */

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		{0, "Shirt", 1 * time.Second},
		{1, "Legos", 1 * time.Second},
		{2, "TV", 5 * time.Second},
		{3, "Bananas", 2 * time.Second},
		{4, "Hat", 1 * time.Second},
		{5, "Phone", 2 * time.Second},
		{6, "Plates", 3 * time.Second},
		{7, "Computer", 5 * time.Second},
		{8, "Pint Glass", 3 * time.Second},
		{9, "Watch", 2 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func PackItems(done <-chan bool, items <-chan Item, workerID int) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Worker #%d: Shipping package no. %d, took %ds to pack\n", workerID, item.ID, item.PackingEffort/time.Second)
			}
		}
		close(packages)
	}()
	return packages
}

func merge(done <-chan bool, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingPackages := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case outgoingPackages <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingPackages)
	}()
	return outgoingPackages
}

func main() {
	fmt.Printf("ENV says GOMAXPROCS: :%s: \n", os.Getenv("GOMAXPROCS"))
	fmt.Printf("runtime says MAXPROCS = %d \n", runtime.NumCPU())

	done := make(chan bool)
	defer close(done)

	start := time.Now()

	items := PrepareItems(done)

	workers := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		workers[i] = PackItems(done, items, i)
	}

	numPackages := 0
	for range merge(done, workers...) {
		numPackages++
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
