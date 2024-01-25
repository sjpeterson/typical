package queue

// Queue is a FIFO ordered collection of items
type Queue[T any] interface {
	Enqueue(item T)
	Dequeue() (T, bool)
	Length() int
	IsEmpty() bool
}

type queue[T any] struct {
	items []T
}

// NewQueue creates a new queue
func NewQueue[T any](items ...T) Queue[T] {
	queueItems := make([]T, 0)
	queueItems = append(queueItems, items...)

	return &queue[T]{
		items: queueItems,
	}
}

// Enqueue adds an item to the end of the queue
func (q *queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue takes an item from the front of the queue
func (q *queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T

		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

// Length returns the number of items in the queue
func (q *queue[T]) Length() int {
	return len(q.items)
}

// IsEmpty returns true if there are no items in the queue
func (q *queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
