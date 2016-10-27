package queue

// Queue
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"runtime"
	"sync/atomic"
)

const sizeBlockDefault int = 1000
const trialLimit int = 20000000

// Message - default element for queue
type Message struct {
	id int
}

// Queue - main struct.
type Queue struct {
	hasp      int32
	db        []Message
	head      int
	tail      int
	sizeQueue int
	sizeBlock int
}

// New - create new queue.
// The optional argument: the initial size of the queue.
func New(args ...int) Queue {
	var sizeBlock int
	if len(args) > 0 {
		sizeBlock = args[0]
	} else {
		sizeBlock = sizeBlockDefault
	}
	q := Queue{
		0, make([]Message, sizeBlock), sizeBlock / 2, sizeBlock / 2, sizeBlock, sizeBlock, // nil,
	}
	q.hasp = 0
	return q
}

// PushTail - Insert element in the tail queue
func (q *Queue) PushTail(n Message) bool {
	if !q.lock() {
		return false
	}
	q.db[q.tail] = n
	q.tail++
	if q.tail >= q.sizeQueue {
		q.db = append(q.db, make([]Message, q.sizeBlock)...)
		q.sizeQueue += q.sizeBlock
	}
	q.hasp = 0
	return true
}

// PushHead - Paste item in the queue head
func (q *Queue) PushHead(n Message) bool {
	if !q.lock() {
		return false
	}
	q.head--
	if q.head == 0 {
		newDb := make([]Message, q.sizeQueue+q.sizeBlock)
		copy(newDb[q.sizeBlock:], q.db)
		q.db = newDb
		q.head += q.sizeBlock
		q.tail += q.sizeBlock
		q.sizeQueue = q.sizeQueue + q.sizeBlock
	}
	q.db[q.head] = n
	q.hasp = 0
	return true
}

// PopHead - Get the first element of the queue
func (q *Queue) PopHead() (Message, bool) {
	var n Message
	if !q.lock() {
		return n, false
	}
	if q.tail == q.head {
		q.hasp = 0
		return n, false
	}
	n, q.db[q.head] = q.db[q.head], Message{}
	q.head++
	if q.head == q.tail {
		q.clean()
	}
	q.hasp = 0
	return n, true
}

// PopTail - Get the item from the queue tail
func (q *Queue) PopTail() (Message, bool) {
	var n Message
	if !q.lock() {
		return n, false
	}
	if q.head == q.tail {
		q.hasp = 0
		return n, false
	}
	q.tail--
	n, q.db[q.tail] = q.db[q.tail], Message{}
	if q.head == q.tail {
		q.clean()
	}
	q.hasp = 0
	return n, true
}

// LenQueue - The number of elements in the queue
func (q *Queue) LenQueue() int {
	q.lock()
	ln := q.tail - q.head
	q.hasp = 0
	return ln
}

// SizeQueue - The size reserved for queue
func (q *Queue) SizeQueue() int {
	q.lock()
	ln := q.sizeQueue
	q.hasp = 0
	return ln
}

// clean - Resetting the queue (not thread-safe, is called only after the lock)
func (q *Queue) clean() {
	q.db = make([]Message, q.sizeBlock)
	q.head = q.sizeBlock / 2
	q.tail = q.sizeBlock / 2
	q.sizeQueue = q.sizeBlock
}

// lock - block queue
func (q *Queue) lock() bool {
	for i := trialLimit; i > 0; i-- {
		if q.hasp == 0 && atomic.CompareAndSwapInt32(&q.hasp, 0, 1) {
			break
		}
		if i == 0 {
			return false
		}
		runtime.Gosched()
	}
	return true
}
