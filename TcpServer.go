package main

import (
	"fmt"
	"net"
)

func handler(c net.Conn, ch chan int) {
	ch <- 11
	c.Write([]byte("OKAY"))
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	ch := make(chan int)
	go logger(ch)
	go server(l, ch)
	select {}
}

func logger(ch chan int) {
	for {
		fmt.Println(<-ch)
	}

}

func server(l net.Listener, ch chan int) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c, ch)
	}
}
