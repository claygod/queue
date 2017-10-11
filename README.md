The library implements the queue with a maximum speed of execution.

[![API documentation](https://godoc.org/github.com/claygod/queue?status.svg)](https://godoc.org/github.com/claygod/queue)

The queue can be used both in style LIFO and a FIFO. Also, urgent messages (with high priority) using FIFO can be put in the queue head.

# Usage

An example of using the Queue Library:
```Go
package main

import (
	"fmt"
	"github.com/claygod/queue"
)

func main() {
	q := New()
	q.PushTail(Message{id: 1})
	q.PushTail(Message{id: 2})
	if m, ok := q.PopHead() {
		fmt.Print("\n id2 = ",q)
	}
	if m, ok := q.PopHead() {
		fmt.Print("\n id1 = ",q)
	}
}
```

### LIFO

```Go
	q := New()
	// in
	q.PushTail(Message{id: 1})
	q.PushTail(Message{id: 2})
	// out
	m, ok := q.PopTail() // -> 2
	m, ok := q.PopTail() // -> 1
```

### FIFO

```Go
	q := New()
	// in
	q.PushTail(Message{id: 1})
	q.PushTail(Message{id: 2})
	// out
	m, ok := q.PopHead() // -> 1
	m, ok := q.PopHead() // -> 2
```

### Priority

```Go
	q := New()
	// in
	q.PushTail(Message{id: 1})
	q.PushTail(Message{id: 2})
	q.PushHead(Message{id: 3}) // immediately in front of the queue
	// out
	m, ok := q.PopHead() // -> 3
	m, ok := q.PopHead() // -> 1
	m, ok := q.PopHead() // -> 2
```

# Perfomance

Operation added to the queue head is more resource-intensive,
so it is necessary to organize the code in the operation of
the library so that the addition was mainly in the tail queue.

Note: in parallel execution speed is decreased. If desired, you can artificially
remove the parallelism in the calling code with a code `runtime.GOMAXPROCS(1)`

```
BenchmarkPushTail-4           	50000000	        31.7 ns/op
BenchmarkPushTailParallel-4   	10000000	        47.7 ns/op
BenchmarkPushHeadLimit-4      	50000000	        29.9 ns/op
BenchmarkPushHead-4           	10000000	      1911.0 ns/op
BenchmarkPopHead-4            	100000000	        15.2 ns/op
BenchmarkPopTail-4            	100000000	        15.0 ns/op
```

# API

Methods:
-  *New* - create new queue.
-  *PushTai*l - insert element in the tail queue
-  *PushHead* - paste item in the queue head
-  *PopHead* - get the first element of the queue
-  *PopTail* - get the item from the queue tail
-  *LenQueue* - the number of elements in the queue
-  *SizeQueue* - the size reserved for queue
-  *clean* - resetting the queue (not thread-safe, is called only after the lock)
-  *lock* - block queue



