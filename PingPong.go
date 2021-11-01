package main

import (
    "fmt"
    "time"
)

func main() {
    playTime := time.Duration(2) //seconds
    next := counter(1)
    numberOfPlayers := 4
    var Ball int
    table := make(chan int)
    for i := 0; i < numberOfPlayers; i++ {
        go player(table, next)
    }

    table <- Ball
    time.Sleep(playTime * time.Second)
    <-table
}

func player(table chan int, next func() int) {
    playerNumber := next()
    for {
        ball := <-table
        ball++
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("Player# %d Hit #%d \n", playerNumber, ball)
        table <- ball
    }
}

func counter(start int) func() int {
    //closure
    count := start - 1
    f := func() int {
        count++
        return count
    }
    return f
}
