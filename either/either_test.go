package either

import "testing"

func TestLeft(t *testing.T) {
	testValue := 450
	wrapped := Left[int, string](testValue)

	if !wrapped.IsLeft() {
		t.Error("IsLeft should be true for Left value")
	}
	if wrapped.IsRight() {
		t.Error("IsRight should be false for Left value")
	}

	unwrapped, ok := wrapped.UnwrapLeft()
	if !ok {
		t.Error("UnwrapLeft should be ok for Left value")
	}
	if unwrapped != testValue {
		t.Errorf("expected unwrapped wrapped to be %d, but got %d", testValue, unwrapped)
	}

	_, ok = wrapped.UnwrapRight()
	if ok {
		t.Error("UnwrapRight should not be ok for Left value")
	}
}

func TestRight(t *testing.T) {
	testValue := "example"
	wrapped := Right[int, string](testValue)

	if wrapped.IsLeft() {
		t.Error("IsLeft should be false for Right value")
	}
	if !wrapped.IsRight() {
		t.Error("IsRight should be true for Right value")
	}

	unwrapped, ok := wrapped.UnwrapRight()
	if !ok {
		t.Error("UnwrapRight should be ok for Right value")
	}
	if unwrapped != testValue {
		t.Errorf("expected unwrapped wrapped to be %s, but got %s", testValue, unwrapped)
	}

	_, ok = wrapped.UnwrapLeft()
	if ok {
		t.Error("UnwrapLeft should not be ok for Right value")
	}
}
