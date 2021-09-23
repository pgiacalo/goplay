//
//  Stack.go

//  This is a stack implementation that is go-routine safe.
//  A stack is a Last-In-First-Out structure (LIFO)
//	This stack implementation can either grow infinitely large or be limited in size.
//	If limited in size, then items at the bottom of thestack are thrown away when the size ie exceeded.
//
//  Phil's Note: this stack implementation is derived from:
//  A FIFO queue created by Hicham Bouabdallah
//  Copyright (c) 2012 SimpleRocket LLC

//	A go-routine safe FIFO (first in first out) data structure.
//  https://github.com/hishboy/gocommons/blob/master/lang/queue.go
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

package main

import (
	"fmt"
	"sync"
)

type stacknode struct {
	data  interface{}
	below *stacknode
	above *stacknode
}

type Stack struct {
	top     *stacknode
	bottom  *stacknode
	count   int
	lock    *sync.Mutex
	maxSize int
}

//	Creates a new pointer to a new stack.
func NewStack() *Stack {
	q := &Stack{}
	q.lock = &sync.Mutex{}
	return q
}

//	Creates a new pointer to a new stack that limits its size
//	If the size exceeds the sizeLimit, the bottom item is removed from the stack.
func NewLimitStack(sizeLimit int) *Stack {
	q := &Stack{}
	q.maxSize = sizeLimit
	q.lock = &sync.Mutex{}
	return q
}

//	Returns the number of elements in the stack (i.e. size/length)
//	go-routine safe.
func (q *Stack) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count
}

//	Pushes a value on top of the stack.
//	Note: this function does mutate the stack.
//	go-routine safe.
func (q *Stack) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &stacknode{data: item}

	if q.bottom == nil {
		q.bottom = n
		q.top = n
	} else {
		n.below = q.top
		n.below.above = n
		q.top = n
	}
	q.count++

	//now keep the stack at the maxSize by pulling items from the bottom of the stack, if needed.
	if q.maxSize > 0 && q.count > q.maxSize {
		doPull(q)
	}
}

//	Returns the value at the top of the stack.
//	i.e. the oldest value in the stack.
//	Note: this function does mutate the stack.
//	go-routine safe.
func (q *Stack) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.top == nil {
		return nil
	}

	n := doPop(q)

	return n.data
}

func doPop(q *Stack) *stacknode {
	if q.top == nil {
		return nil
	}

	n := q.top
	q.top = n.below
	if q.top != nil {
		q.top.above = nil
	}

	if q.top == nil {
		q.bottom = nil
	}
	q.count--

	return n
}

//	Returns a read value at the top of the stack.
//	i.e. the newest value in the stack.
//	Note: this function does NOT mutate the stack.
//	go-routine safe.
func (q *Stack) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.top
	if n == nil {
		return nil
	}

	return n.data
}

//	Returns the value at the bottom of the stack.
//	i.e. the oldest value in the stack.
//	Note: this function does mutate the stack.
//	go-routine safe.
func (q *Stack) Pull() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.top == nil {
		return nil
	}

	n := doPull(q)

	return n.data
}

func doPull(q *Stack) *stacknode {
	if q.top == nil {
		return nil
	}

	n := q.bottom
	q.bottom = n.above
	if q.bottom != nil {
		q.bottom.below = nil
	}

	if q.top == nil {
		q.bottom = nil
	}
	q.count--

	return n
}

// Empties the Stack
// go-routine safe.
func (q *Stack) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	for {
		if q.top == nil {
			return
		}

		n := q.top
		q.top = n.below

		if q.top == nil {
			q.bottom = nil
		}
		q.count--
	}
}

func main() {
	q := NewLimitStack(3)

	q.Push("A")
	q.Push("B")
	q.Push("C")
	q.Push("D")
	q.Push("E")
	fmt.Println(q.Len())

	fmt.Println(q.Pop())

	q.Pull()

	q.Push("R")

	popPrint(q)

	// fmt.Println(q.Pop())

	q.Clear()

	q.Push("S")
	q.Push("T")
	q.Push("U")
	q.Push("V")
	// fmt.Println(q.Peek())
	q.Pull()

	popPrint(q)

}

func popPrint(q *Stack) {
	for {
		n := q.Pop()
		if n == nil {
			break
		}
		fmt.Println(n)
	}
}
