package stack

// Stack is a LIFO ordered collection of items.
type Stack[T any] interface {
	Push(item T)
	Pop() (T, bool)
	Size() int
	IsEmpty() bool
}

type stack[T any] struct {
	items []T
}

// NewStack creates a new stack.
func NewStack[T any](items ...T) Stack[T] {
	stackItems := make([]T, 0)
	stackItems = append(stackItems, items...)

	return &stack[T]{
		items: stackItems,
	}
}

// Push puts a new item at the top of the stack.
func (st *stack[T]) Push(item T) {
	st.items = append(st.items, item)
}

// Pop takes an item from the top of the stack.
func (st *stack[T]) Pop() (T, bool) {
	if len(st.items) == 0 {
		var zero T

		return zero, false
	}

	lastIdx := len(st.items) - 1
	item := st.items[lastIdx]
	st.items = st.items[:lastIdx]

	return item, true
}

// Size returns the number of items in the stack.
func (st *stack[T]) Size() int {
	return len(st.items)
}

// IsEmpty returns true if there are no items in the stack.
func (st *stack[T]) IsEmpty() bool {
	return len(st.items) == 0
}
