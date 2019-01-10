package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	for {
		ticker := <-stockChannel
		io.WriteString(w, fmt.Sprintf("%d", ticker))
	}
}

// setup a channel for delivering stocks
var stockChannel = make(chan string, 5)

func SendTicker() {
	var tick string = "Apple"
	for {
		stockChannel <- tick
		//tick += 1
		time.Sleep(3 * 1e9)
	}
}

func main() {
	http.HandleFunc("/", handler)

	go SendTicker()
	err := http.ListenAndServe("127.0.0.1:5900", nil)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
