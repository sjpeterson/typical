package tuples

import "testing"

func TestNewPair(t *testing.T) {
	first := "test"
	second := 45

	testPair := NewPair(first, second)
	if testPair.First != first {
		t.Errorf("expected first member to be %q, but got %q", first, testPair.First)
	}
	if testPair.Second != second {
		t.Errorf("expected second member to be %v, but got %v", second, testPair.Second)
	}
}

func TestDestructurePair(t *testing.T) {
	testPair := Pair[string, int]{
		First:  "test",
		Second: 45,
	}
	first, second := testPair.Destructure()

	if first != testPair.First {
		t.Errorf("expected first value to be %q, but got %q", testPair.First, first)
	}
	if second != testPair.Second {
		t.Errorf("expected second value to be %v, but got %v", testPair.Second, second)
	}

}

func TestNewTup3(t *testing.T) {
	first := "test"
	second := 45
	third := -17.3

	testTup := NewTup3(first, second, third)
	if testTup.First != first {
		t.Errorf("expected first member to be %q, but got %q", first, testTup.First)
	}
	if testTup.Second != second {
		t.Errorf("expected second member to be %v, but got %v", second, testTup.Second)
	}
	if testTup.Third != third {
		t.Errorf("expected third member to be %v, but got %v", third, testTup.Third)
	}
}

func TestDestructureTup3(t *testing.T) {
	testTup := Tup3[string, int, float64]{
		First:  "test",
		Second: 45,
		Third:  -17.3,
	}
	first, second, third := testTup.Destructure()

	if first != testTup.First {
		t.Errorf("expected first value to be %q, but got %q", testTup.First, first)
	}
	if second != testTup.Second {
		t.Errorf("expected second value to be %v, but got %v", testTup.Second, second)
	}
	if third != testTup.Third {
		t.Errorf("expected third value to be %v, but got %v", testTup.Third, third)
	}

}

func TestNewTup4(t *testing.T) {
	type TestStruct struct {
		foo string
		bar uint16
	}

	first := "test"
	second := 45
	third := -17.3
	fourth := TestStruct{
		foo: "more test",
		bar: 4,
	}

	testTup := NewTup4(first, second, third, fourth)
	if testTup.First != first {
		t.Errorf("expected first member to be %q, but got %q", first, testTup.First)
	}
	if testTup.Second != second {
		t.Errorf("expected second member to be %v, but got %v", second, testTup.Second)
	}
	if testTup.Third != third {
		t.Errorf("expected third member to be %v, but got %v", third, testTup.Third)
	}
	if testTup.Fourth != fourth {
		t.Errorf("expected fourth member to be %v, but got %v", fourth, testTup.Fourth)
	}
}

func TestDestructureTup4(t *testing.T) {
	type TestStruct struct {
		foo string
		bar uint16
	}

	testTup := Tup4[string, int, float64, TestStruct]{
		First:  "test",
		Second: 45,
		Third:  -17.3,
		Fourth: TestStruct{
			foo: "more test",
			bar: 4,
		},
	}
	first, second, third, fourth := testTup.Destructure()

	if first != testTup.First {
		t.Errorf("expected first value to be %q, but got %q", testTup.First, first)
	}
	if second != testTup.Second {
		t.Errorf("expected second value to be %v, but got %v", testTup.Second, second)
	}
	if third != testTup.Third {
		t.Errorf("expected third value to be %v, but got %v", testTup.Third, third)
	}
	if fourth != testTup.Fourth {
		t.Errorf("expected fourth value to be %v, but got %v", testTup.Fourth, fourth)
	}
}
