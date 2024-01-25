package stack

import "testing"

func TestNewStack(t *testing.T) {
	emptyStack := NewStack[int]()
	if size := emptyStack.Size(); size != 0 {
		t.Errorf("expected empty stack size to be 0, but got %d", size)
	}
	singletonStack := NewStack[int](1)
	if size := singletonStack.Size(); size != 1 {
		t.Errorf("expected singleton stack size to be 1, but got %d", size)
	}
	testStack := NewStack[int](5, 8, 13)
	if size := testStack.Size(); size != 3 {
		t.Errorf("expected stack size to be 3, but got %d", size)
	}
}

func TestStack_Push(t *testing.T) {
	testStack := NewStack[int]()
	if size := testStack.Size(); size != 0 {
		t.Errorf("expected empty stack size to be 0, but got %d", size)
	}
	testStack.Push(5)
	if size := testStack.Size(); size != 1 {
		t.Errorf("expected stack size after first push to be 1, but got %d", size)
	}
	testStack.Push(3)
	if size := testStack.Size(); size != 2 {
		t.Errorf("expected stack size after second push to be 2, but got %d", size)
	}
}

func TestStack_Pop(t *testing.T) {
	testStack := NewStack(6, 3, 2)
	first, ok := testStack.Pop()
	if !ok {
		t.Error("failed to pop item from populated stack")
	}
	second, ok := testStack.Pop()
	if !ok {
		t.Error("failed to pop item from populated stack")
	}
	third, ok := testStack.Pop()
	if !ok {
		t.Error("failed to pop item from populated stack")
	}
	_, ok = testStack.Pop()
	if ok {
		t.Error("succeeded popping item from empty stack")
	}

	if first != 2 {
		t.Errorf("expected first item popped to be 2, but got %d", first)
	}
	if second != 3 {
		t.Errorf("expected second item popped to be 3, but got %d", second)
	}
	if third != 6 {
		t.Errorf("expected third item popped to be 6, but got %d", third)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	testStack := NewStack[int]()
	if !testStack.IsEmpty() {
		t.Error("expected IsEmpty to return true for empty stack, but got false")
	}
	testStack.Push(6)
	if testStack.IsEmpty() {
		t.Error("expected IsEmpty to return false for non-empty stack, but got true")
	}
	_, _ = testStack.Pop()
	if !testStack.IsEmpty() {
		t.Error("expected IsEmpty to return true for emptied stack, but got false")
	}
}

func TestStack_LIFO(t *testing.T) {
	testItems := []int{6, 3, 2}
	testStack := NewStack[int]()

	for _, item := range testItems {
		testStack.Push(item)
	}

	for k := range testItems {
		item, ok := testStack.Pop()
		if !ok {
			t.Error("failed to pop item from populated stack")
		}
		expected := testItems[len(testItems)-k-1]
		if item != expected {
			t.Errorf("expected item %d popped to be %d, but got %d", k, expected, item)
		}
	}

}
