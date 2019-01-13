// This package holds the code from
// http://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/
// updated to Go 1 standard. In particular, it's now OK to pass around
// by-value objects containing private fields, and we don't need to use
// semicolons.
package main

import (
	"fmt"
	"time"
)

type broadcast struct {
	c chan broadcast
	v interface{}
}

// Broadcaster allows
type Broadcaster struct {
	listenc chan chan (chan broadcast)
	sendc   chan<- interface{}
}

// Receiver can be used to wait for a broadcast value.
type Receiver struct {
	c chan broadcast
}

// NewBroadcaster returns a new broadcaster object.
func NewBroadcaster() Broadcaster {
	listenc := make(chan (chan (chan broadcast)))
	sendc := make(chan interface{})
	go func() {
		currc := make(chan broadcast, 1)
		for {
			select {
			case v := <-sendc:
				if v == nil {
					currc <- broadcast{}
					return
				}
				c := make(chan broadcast, 1)
				b := broadcast{c: c, v: v}
				currc <- b
				currc = c
			case r := <-listenc:
				r <- currc
			}
		}
	}()
	return Broadcaster{
		listenc: listenc,
		sendc:   sendc,
	}
}

// Listen starts returns a Receiver that
// listens to all broadcast values.
func (b Broadcaster) Listen() Receiver {
	c := make(chan chan broadcast, 0)
	b.listenc <- c
	return Receiver{<-c}
}

// Write broadcasts a a value to all listeners.
func (b Broadcaster) Write(v interface{}) {
	b.sendc <- v
}

// Read reads a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
	b := <-r.c
	v := b.v
	r.c <- b
	r.c = b.c
	return v
}

//--------- testing code ----------

type message struct {
	id int
	name string
}

var b = NewBroadcaster();

func listen(id int, r Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
//		go listen(r);
		fmt.Printf("listener id:%v, received message:%v\n", id, v);
	}
}

func main() {
	receiver1 := b.Listen();
	receiver2 := b.Listen();

	//start two listeners to listen for broadcasts from broadcaster b
	go listen(1, receiver1);
	go listen(2, receiver2);

	for i := 0; i  < 10; i++ {
		m := message{i, "Hello World"}
		b.Write(m);
		time.Sleep(250 * time.Millisecond)
	}
	b.Write(nil);

	time.Sleep(3 * 1e9);
}
