package main

// references for this code:
// https://play.golang.org/p/HCbY04zIg3
//
// http://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/
// updated to Go 1 standard. In particular, it's now OK to pass around
// by-value objects containing private fields, and we don't need to use
// semicolons.

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
	cc    chan broadcast
	sendc chan<- interface{}
}

// Receiver can be used to wait for a broadcast value.
type Receiver struct {
	c chan broadcast
}

// NewBroadcaster returns a new broadcaster object.
func NewBroadcaster() Broadcaster {
	cc := make(chan broadcast, 1)
	sendc := make(chan interface{})
	b := Broadcaster{
		sendc: sendc,
		cc:    cc,
	}

	go func() {
		for {
			select {
			case v := <-sendc:
				if v == nil {
					b.cc <- broadcast{}
					return
				}
				c := make(chan broadcast, 1)
				newb := broadcast{c: c, v: v}
				b.cc <- newb
				b.cc = c
			}
		}
	}()

	return b
}

// Listen starts returns a Receiver that
// listens to all broadcast values.
func (b Broadcaster) Listen() Receiver {
	return Receiver{b.cc}
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

type messageA struct {
	id int
	name string
}

type messageB struct {
	id int
	value float64
	msg messageA
}

func listen(id int, r Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		//		go listen(r);
		fmt.Printf("listener id:%v, received message:%v\n", id, v);
	}
}

// setup 2 separate broadcasters (i.e., two separate publisher topics)
var a = NewBroadcaster();
var b = NewBroadcaster();

func main() {
	//create 2 receivers (i.e., subscribers) to publisher "a"
	receiverA1 := a.Listen();
	receiverA2 := a.Listen();

	//create 2 receivers (i.e., subscribers) to publisher "b"
	receiverB1 := b.Listen();
	receiverB2 := b.Listen();

	//start two listeners to listen for broadcasts from broadcaster a
	go listen(1, receiverA1);
	go listen(2, receiverA2);

	//start two listeners to listen for broadcasts from broadcaster b
	go listen(3, receiverB1);
	go listen(4, receiverB2);

	//publish messages via publisher a
	go sendA()
	//publish messages via publisher b
	go sendB()

	//keep main alive so the work can get done
	time.Sleep(3 * 1e9);
}

func sendA(){
	for i := 0; i  < 10; i++ {
		msgA := messageA{i, "Msg A name"}
		a.Write(msgA);
		time.Sleep(250 * time.Millisecond)
	}
	a.Write(nil);
}

func sendB(){
	for i := 0; i  < 10; i++ {
		msgA := messageA{i, "Embedded msg A name"}
		msgB := messageB{i, 3.1415, msgA}
		b.Write(msgB);
		time.Sleep(250 * time.Millisecond)
	}
	b.Write(nil);
}