package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	emptyQueue := NewQueue[int]()
	if length := emptyQueue.Length(); length != 0 {
		t.Errorf("expected empty queue length to be 0, but got %d", length)
	}
	singletonQueue := NewQueue[int](1)
	if length := singletonQueue.Length(); length != 1 {
		t.Errorf("expected singleton queue length to be 1, but got %d", length)
	}
	testQueue := NewQueue[int](5, 8, 13)
	if length := testQueue.Length(); length != 3 {
		t.Errorf("expected queue length to be 3, but got %d", length)
	}
}

func TestQueue_Enqueue(t *testing.T) {
	testQueue := NewQueue[int]()
	if length := testQueue.Length(); length != 0 {
		t.Errorf("expected empty queue length to be 0, but got %d", length)
	}
	testQueue.Enqueue(5)
	if length := testQueue.Length(); length != 1 {
		t.Errorf("expected queue length after adding the first item to be 1, but got %d", length)
	}
	testQueue.Enqueue(3)
	if length := testQueue.Length(); length != 2 {
		t.Errorf("expected queue length after adding the second item to be 2, but got %d", length)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	testQueue := NewQueue(6, 3, 2)
	first, ok := testQueue.Dequeue()
	if !ok {
		t.Error("failed to dequeue item from populated queue")
	}
	second, ok := testQueue.Dequeue()
	if !ok {
		t.Error("failed to dequeue item from populated queue")
	}
	third, ok := testQueue.Dequeue()
	if !ok {
		t.Error("failed to dequeue item from populated queue")
	}
	_, ok = testQueue.Dequeue()
	if ok {
		t.Error("succeeded dequeue item from empty queue")
	}

	if first != 6 {
		t.Errorf("expected first item dequeued to be 6, but got %d", first)
	}
	if second != 3 {
		t.Errorf("expected second item dequeued to be 3, but got %d", second)
	}
	if third != 2 {
		t.Errorf("expected third item dequeued to be 2, but got %d", third)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	testQueue := NewQueue[int]()
	if !testQueue.IsEmpty() {
		t.Error("expected IsEmpty to return true for empty queue, but got false")
	}
	testQueue.Enqueue(6)
	if testQueue.IsEmpty() {
		t.Error("expected IsEmpty to return false for non-empty queue, but got true")
	}
	_, _ = testQueue.Dequeue()
	if !testQueue.IsEmpty() {
		t.Error("expected IsEmpty to return true for emptied queue, but got false")
	}
}

func TestQueue_FIFO(t *testing.T) {
	testItems := []int{6, 3, 2}
	testQueue := NewQueue[int]()
	for _, item := range testItems {
		testQueue.Enqueue(item)
	}

	for k, expected := range testItems {
		item, ok := testQueue.Dequeue()
		if !ok {
			t.Error("failed to dequeue item from populated queue")
		}
		if item != expected {
			t.Errorf("expected item %d dequeued to be %d, but got %d", k, expected, item)
		}
	}
}
