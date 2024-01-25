package maybe

import (
	"testing"
)

func TestSomeIsSome(t *testing.T) {
	someInt := Some(6)
	someString := Some("test")

	if !someInt.IsSome() {
		t.Error("Some(6).IsSome() is false")
	}

	if !someString.IsSome() {
		t.Error("Some(\"test\").IsSome() is false")
	}
}

func TestSomeIsNotNone(t *testing.T) {
	someInt := Some(6)
	someString := Some("test")

	if someInt.IsNone() {
		t.Error("Some(6).IsNone() is true")
	}

	if someString.IsNone() {
		t.Error("Some(\"test\").IsNone() is true")
	}
}

func TestNoneIsNotSome(t *testing.T) {
	noInt := None[int]()
	noString := None[string]()

	if noInt.IsSome() {
		t.Error("None[int]().IsSome() is true")
	}

	if noString.IsSome() {
		t.Error("None[string]().IsSome() is true")
	}
}

func TestNoneIsNone(t *testing.T) {
	noInt := None[int]()
	noString := None[string]()

	if !noInt.IsNone() {
		t.Error("None[int]().IsNone() is false")
	}

	if !noString.IsNone() {
		t.Error("None[string]().IsNone() is false")
	}
}

func TestSome_Unwrap(t *testing.T) {
	someInt := Some(6)
	someString := Some("test")
	noInt := None[int]()
	noString := None[string]()

	someIntValue, someIntOk := someInt.Unwrap()
	someStringValue, someStringOk := someString.Unwrap()
	_, noIntOk := noInt.Unwrap()
	_, noStringOk := noString.Unwrap()

	if someIntValue != 6 {
		t.Errorf("Unwrapping Some(6) returned %d", someIntValue)
	}

	if !someIntOk {
		t.Error("Unwrapping Some(6) was not ok")
	}

	if someStringValue != "test" {
		t.Errorf("Unwrapping Some(\"test\") returned \"%s\"", someStringValue)
	}

	if !someStringOk {
		t.Error("Unwrapping Some(\"test\") was not ok")
	}

	if noIntOk {
		t.Error("Unwrapping None[int]() was ok")
	}

	if noStringOk {
		t.Error("Unwrapping None[string]() was ok")
	}
}

