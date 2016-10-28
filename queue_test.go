package queue

// Queue
// Test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

func TestLenQueue(t *testing.T) {
	q := New(10)
	q.PushTail(Message{id: 1})
	q.PushTail(Message{id: 2})
	if q.LenQueue() != 2 {
		t.Error("LenQueue 2 != ", q.LenQueue())
	}
}

func TestSizeQueue10(t *testing.T) {
	q := New(10)
	q.PushTail(Message{id: 1})
	if q.SizeQueue() != 10 {
		t.Error("SizeQueue 10 != ", q.LenQueue())
	}
}

func TestSizeQueue20(t *testing.T) {
	q := New(10)
	for i := 0; i < 5; i++ {
		q.PushTail(Message{id: 1})
	}
	if q.SizeQueue() != 20 {
		t.Error("SizeQueue 20 != ", q.LenQueue())
	}
}

func TestClear(t *testing.T) {
	q := New(10)
	for i := 0; i < 30; i++ {
		q.PushTail(Message{id: 1})
	}
	for i := 0; i < 30; i++ {
		q.PopTail()
	}
	if q.SizeQueue() != 10 {
		t.Error("Clear not correct! SizeQueue 10 != ", q.SizeQueue())
	}
}

func TestDatabase(t *testing.T) {
	q := New(10)
	for i := 0; i < 100; i++ {
		q.PushTail(Message{id: i})
	}
	for i := 99; i >= 0; i-- {
		m, _ := q.PopTail()
		if m.id != i {
			t.Error("Error in database", i, " != ", m.id)
		}
	}
}
func TestPopHead(t *testing.T) {
	q := New(10)
	q.PushTail(Message{id: 100})
	q.PushHead(Message{id: 200})
	q.PushTail(Message{id: 300})

	if m, _ := q.PopHead(); m.id != 200 {
		t.Error("Error PopHead: 200 != ", m.id)
	}
	if m, _ := q.PopHead(); m.id != 100 {
		t.Error("Error PopHead: 100 != ", m.id)
	}
	if m, _ := q.PopHead(); m.id != 300 {
		t.Error("Error PopHead: 300 != ", m.id)
	}
}

func TestPopTail(t *testing.T) {
	q := New(10)
	q.PushTail(Message{id: 100})
	q.PushHead(Message{id: 200})
	q.PushTail(Message{id: 300})

	if m, _ := q.PopTail(); m.id != 300 {
		t.Error("Error PopTail: 300 != ", m.id)
	}
	if m, _ := q.PopTail(); m.id != 100 {
		t.Error("Error PopTail: 100 != ", m.id)
	}
	if m, _ := q.PopTail(); m.id != 200 {
		t.Error("Error PopTail: 200 != ", m.id)
	}
}
