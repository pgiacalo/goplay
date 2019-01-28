package main

// This is an implementation of a Publish/Subscribe pattern.
// -- PubSub.go uses struct and field names that are more intuitive (to Phil) than PublishSubscribe.go --
// It enables messages to be published to multiple subscribers.
// Note that each receiver gets its own COPY of the original message.
//
// See the test code at the bottom of this file
//
// references for this code:
// Source code:
// https://play.golang.org/p/HCbY04zIg3
//
// Discussion:
// http://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-publicationing-values-in-go-with-linked-channels/
// updated to Go 1 standard. In particular, it's now OK to pass around
// by-value objects containing private fields, and we don't need to use
// semicolons.

import (
	"fmt"
	"time"
)

type message struct {
	//m holds the value of the message being published
	m interface{}
	//the message struct also has a channel member (c) that holds the publisher's channel.
	//this is required in order to pass the message along to the next subscriber, after the message is received.
	c chan message
}

// Publisher allows
type Publisher struct {
	//inchan is the channel used to receive messages to be published
	inchan chan<- interface{}
	//outchan is the channel used to deliver messages to the subscribers
	outchan chan message
}

// A Subscriber waits and receives message values being published.
type Subscriber struct {
	msgchan chan message
}

// NewPublisher starts and returns a new Publisher object.
// The returned Publisher is waiting for messages to arrive via its inchan
// and will publish those messages via its outchan to all subscribers.
func NewPublisher() Publisher {
	inchan := make(chan interface{})
	outchan := make(chan message, 1)
	p := Publisher{
		inchan:  inchan,
		outchan: outchan,
	}

	go func() {
		for {
			select {
			case v := <-inchan:
				if v == nil {
					p.outchan <- message{}
					return
				}
				c := make(chan message, 1)
				newmsg := message{c: c, m: v}
				p.outchan <- newmsg
				p.outchan = c
			}
		}
	}()

	return p
}

// Subscribe returns a Subscriber that
// listens to all message values.
func (p Publisher) Subscribe() Subscriber {
	return Subscriber{p.outchan}
}

// Publish publications to all listeners.
func (p Publisher) Publish(v interface{}) {
	//put the message (v) into the publisher's inchan
	p.inchan <- v
}

// Read reads a message value that has been published, waiting until one is available if necessary.
// Note: the Read() function also resubmits the message to the channel, so that other subscribers will also receive it.
func (s *Subscriber) Read() interface{} {
	//get the message object from the subscriber's channel
	msg := <-s.msgchan
	//get the actual message from the message object
	m := msg.m
	//put the msg back into the same channel so that any other subscribers will also receive it
	s.msgchan <- msg
	//set the subscriber's msg channel to the message's channel (which is the Publisher's channel)
	s.msgchan = msg.c
	//now return the message to the subscriber
	return m
}

//--------- test code ----------

//a message struct that will be published
type messageA struct {
	id   int
	name string
}

//a message struct that will be published
type messageB struct {
	id   int
	name string
	msg  messageA
}

//implementation of the listener that waits for messages to arrive via the given receiver (the id is just for debugging)
func listen(id int, s Subscriber) {
	for v := s.Read(); v != nil; v = s.Read() {
		fmt.Printf("listener id:%v, received message:%v\n", id, v)
		//collect all the messages in one array -- just to double check what is being received
		msgCollection[msgCounter] = v
		msgCounter++
	}
}

var msgsToSend = 5
var msgCounter = 0
var msgCollection []interface{} = make([]interface{}, 4*msgsToSend) //4 subscribers

// setup 2 separate publicationers (i.e., two separate publisher topics)
var pa = NewPublisher()
var pb = NewPublisher()

//create 2 subscribers to publisher "pa"
var subscriberA1 = pa.Subscribe()
var subscriberA2 = pa.Subscribe()

//create 2 subscribers to publisher "pb"
var subscriberB1 = pb.Subscribe()
var subscriberB2 = pb.Subscribe()

func main() {

	//start two listeners -- to listen for publications from publicationer a
	go listen(1, subscriberA1)
	go listen(2, subscriberA2)

	//start two listeners -- to listen for publications from publicationer b
	go listen(3, subscriberB1)
	go listen(4, subscriberB2)

	//publish some simple messages via publisher a
	go sendA(msgsToSend)

	//publish some nested messages via publisher b
	go sendB(msgsToSend)

	//keep main alive so the work can get done
	time.Sleep(3 * 1e9)

	//Finally, print out all of the messages received so we can see exactly what was received.
	//Note that each message has its own address, indicating that the messages are copies of the original.
	for i := 0; i < len(msgCollection); i++ {
		fmt.Printf("msgCollection[%v]: %v , addr: %v\n", i, msgCollection[i], &msgCollection[i])
	}
}

func sendA(qty int) {
	for i := 0; i < qty; i++ {
		msgA := messageA{i, "Msg A"}
		pa.Publish(msgA)
		time.Sleep(250 * time.Millisecond)
	}
	//writing nil closes the channel
	pa.Publish(nil)
}

func sendB(qty int) {
	for i := 0; i < qty; i++ {
		msgA := messageA{i, "Embedded msg A"}
		msgB := messageB{i, "Msg B", msgA}
		pb.Publish(msgB)
		time.Sleep(250 * time.Millisecond)
	}
	//writing nil closes the channel
	pb.Publish(nil)
}
