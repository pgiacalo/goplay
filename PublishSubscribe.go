package main

// This is an implementation of a Publish/Subscribe pattern.
// It enables messages to be published to multiple receivers (subscribers).
// Note that each receiver gets its own COPY of the original message.
//
// See the test code at the bottom of this file
//
// references for this code:
// Source code:
// https://play.golang.org/p/HCbY04zIg3
//
// Discussion:
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

// Listen returns a Receiver that
// listens to all broadcast values.
func (b Broadcaster) Listen() Receiver {
	return Receiver{b.cc}
}

// Write broadcasts a value to all listeners.
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


//--------- test code ----------

//a message struct that will be published
type messageA struct {
	id int
	name string
}

//a message struct that will be published
type messageB struct {
	id int
	name string
	msg messageA
}

//implementation of the listener that waits for messages to arrive via the given receiver (the id is just for debugging)
func listen(id int, r Receiver) {
	for v := r.Read(); v != nil; v = r.Read() {
		fmt.Printf("listener id:%v, received message:%v\n", id, v);
		//collect all the messages in one array -- just to double check what is being received
		msgCollection[msgCounter] = v
		msgCounter++
	}
}

var msgsToSend = 5
var msgCounter = 0
var msgCollection []interface{} = make([]interface{}, 4*msgsToSend) //4 subscribers

// setup 2 separate broadcasters (i.e., two separate publisher topics)
var a = NewBroadcaster()
var b = NewBroadcaster()

//create 2 receivers (i.e., subscribers) to publisher "a"
var receiverA1 = a.Listen()
var receiverA2 = a.Listen()

//create 2 receivers (i.e., subscribers) to publisher "b"
var receiverB1 = b.Listen()
var receiverB2 = b.Listen()

func main() {

	//start two listeners -- to listen for broadcasts from broadcaster a
	go listen(1, receiverA1);
	go listen(2, receiverA2);

	//start two listeners -- to listen for broadcasts from broadcaster b
	go listen(3, receiverB1);
	go listen(4, receiverB2);

	//publish some messages via publisher a
	go sendA(msgsToSend)
	//publish some messages via publisher b
	go sendB(msgsToSend)

	//keep main alive so the work can get done
	time.Sleep(3 * 1e9);

	//Finally, print out all of the messages received so we can see exactly what was received.
	//Note that each message has its own address, indicating that the messages are copies of the original.
	for i:=0; i<len(msgCollection); i++{
		fmt.Printf("msgCollection[%v]: %v, addr: %v\n", i, msgCollection[i], &msgCollection[i])
	}
}

func sendA(qty int){
	for i := 0; i  < qty; i++ {
		msgA := messageA{i, "Msg A"}
		a.Write(msgA);
		time.Sleep(250 * time.Millisecond)
	}
	//writing nil closes the
	a.Write(nil);
}

func sendB(qty int){
	for i := 0; i  < qty; i++ {
		msgA := messageA{i, "Embedded msg A"}
		msgB := messageB{i, "Msg B", msgA}
		b.Write(msgB);
		time.Sleep(250 * time.Millisecond)
	}
	b.Write(nil);
}