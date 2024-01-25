package set

import (
	"cmp"
	"sort"
	"testing"
)

func TestSet_Contains(t *testing.T) {
	testSet := NewSet(2, 3, 5, 7, 11, 13)

	if !testSet.Contains(7) {
		t.Error("NewSet(2, 3, 5, 7, 11, 13).Contains(7) is false")
	}

	if testSet.Contains(8) {
		t.Error("NewSet(2, 3, 5, 7, 11, 13).Contains(8) is true")
	}
}

func TestSet_DoesNotContain(t *testing.T) {
	testSet := NewSet(2, 3, 5, 7, 11, 13)

	if testSet.DoesNotContain(7) {
		t.Error("NewSet(2, 3, 5, 7, 11, 13).DoesNotContain(7) is true")
	}

	if !testSet.DoesNotContain(8) {
		t.Error("NewSet(2, 3, 5, 7, 11, 13).DoesNotContain(8) is false")
	}
}

func TestNewSet(_ *testing.T) {
	_ = NewSet[int]()
	_ = NewSet[string]()
	_ = NewSet(5, 7, 3)
	_ = NewSet("foo", "bar", "baz")
}

func TestSet_Cardinality(t *testing.T) {
	emptySet := NewSet[string]()
	singletonSet := NewSet(5)
	primesBelowFifteen := NewSet(2, 3, 5, 7, 11, 13)

	if s := emptySet.Cardinality(); s != 0 {
		t.Errorf("expected cardinality of empty set to be 0, but got %d", s)
	}

	if s := singletonSet.Cardinality(); s != 1 {
		t.Errorf("expected cardinality of singleton set to be 1, but got %d", s)
	}

	if s := primesBelowFifteen.Cardinality(); s != 6 {
		t.Errorf("expected cardinality of the set of primes below fifteen to be 6, but got %d", s)
	}
}

func TestSet_Add(t *testing.T) {
	testSet := NewSet[int]()

	testSet.Add(7)
	if !testSet.Contains(7) {
		t.Error("set does not contain added element")
	}
	if s := testSet.Cardinality(); s != 1 {
		t.Errorf("expected cardinality to be 1 after adding an element, but got %d", s)
	}

	testSet.Add(8)
	if !testSet.Contains(7) {
		t.Error("set does not contain previously added element")
	}
	if !testSet.Contains(8) {
		t.Error("set does not contain newly added element")
	}
	if s := testSet.Cardinality(); s != 2 {
		t.Errorf("expected cardinality to be 2 after adding a second element, but got %d", s)
	}

	testSet.Add(7)
	if !testSet.Contains(7) {
		t.Error("set does not contain re-added element")
	}
	if s := testSet.Cardinality(); s != 2 {
		t.Errorf("expected cardinality to be 2 after adding an already included element, but got %d", s)
	}
}

func TestSet_Elements(t *testing.T) {
	emptySet := NewSet[string]()
	singletonSet := NewSet(5)
	primesBelowFifteen := NewSet(2, 3, 5, 7, 11, 13)

	if elements := emptySet.Elements(); !areSetEqual(elements, []string{}) {
		t.Errorf("expected elements of the empty set of strings to be an empty slice, but got %v", elements)
	}

	if elements := singletonSet.Elements(); !areSetEqual(elements, []int{5}) {
		t.Errorf("expected elements of the singleton set {5} to be []int{5}, but got %v", elements)
	}

	if elements := primesBelowFifteen.Elements(); !areSetEqual(elements, []int{2, 3, 5, 7, 11, 13}) {
		t.Errorf("expected elements of the set of primes below fifteen to be []int{2, 3, 5, 7, 11, 13}, but got %v", elements)
	}
}

func TestSet_Union(t *testing.T) {
	a := NewSet(2, 4, 6, 8)
	b := NewSet(1, 2, 3, 4, 5)

	a.Union(b)

	if elements := a.Elements(); !areSetEqual(elements, []int{1, 2, 3, 4, 5, 6, 8}) {
		t.Errorf("expected a u b to be {1, 2, 3, 4, 5, 6, 8}, but got %v", elements)
	}

	if !areSetEqual(b.Elements(), []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected b to be unmodified")
	}
}

func TestSet_Intersection(t *testing.T) {
	a := NewSet(2, 4, 6, 8)
	b := NewSet(1, 2, 3, 4, 5)

	a.Intersection(b)

	if elements := a.Elements(); !areSetEqual(elements, []int{2, 4}) {
		t.Errorf("expected a to be {2, 4}, but got %v", elements)
	}

	if !areSetEqual(b.Elements(), []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected b to be unmodified")
	}
}

func TestSet_Difference(t *testing.T) {
	a := NewSet(2, 4, 6, 8)
	b := NewSet(1, 2, 3, 4, 5)

	a.Difference(b)

	if elements := a.Elements(); !areSetEqual(elements, []int{6, 8}) {
		t.Errorf("expected a to be {6, 8}, but got %v", elements)
	}

	if !areSetEqual(b.Elements(), []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected b to be unmodified")
	}
}

func TestSet_SymmetricDifference(t *testing.T) {
	a := NewSet(2, 4, 6, 8)
	b := NewSet(1, 2, 3, 4, 5)

	a.SymmetricDifference(b)

	if elements := a.Elements(); !areSetEqual(elements, []int{1, 3, 5, 6, 8}) {
		t.Errorf("expected a to be {1, 3, 5, 6, 8}, but got %v", elements)
	}

	if !areSetEqual(b.Elements(), []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected b to be unmodified")
	}
}

func TestUnion(t *testing.T) {
	if elements := Union[int]().Elements(); !areSetEqual(elements, []int{}) {
		t.Errorf("expected the union of no sets to be the empty set, but got %v", elements)
	}
	if elements := Union(NewSet(1, 2, 3, 4)).Elements(); !areSetEqual(elements, []int{1, 2, 3, 4}) {
		t.Errorf("expected the union of a single set {1, 2, 3, 4} to be itself, but got %v", elements)
	}

	a := NewSet(2, 4, 6, 8)
	b := NewSet(1, 2, 3, 4, 5)
	c := NewSet(9, 10)

	unionOfAAndB := Union(a, b)
	if elements := unionOfAAndB.Elements(); !areSetEqual(elements, []int{1, 2, 3, 4, 5, 6, 8}) {
		t.Errorf("expected the union of {2, 4, 6, 8} and {1, 2, 3, 4, 5} to be {1, 2, 3, 4, 5, 6, 8}, but got %v", elements)
	}
	if elements := a.Elements(); !areSetEqual(elements, []int{2, 4, 6, 8}) {
		t.Errorf("expected a to be unmodified, but got %v", elements)
	}
	if elements := b.Elements(); !areSetEqual(elements, []int{1, 2, 3, 4, 5}) {
		t.Errorf("expected b to be unmodified, but got %v", elements)
	}

	if elements := Union(a, b, c).Elements(); !areSetEqual(elements, []int{1, 2, 3, 4, 5, 6, 8, 9, 10}) {
		t.Errorf("expected the union of {2, 4, 6, 8}, {1, 2, 3, 4, 5}, and {9, 10} to be {1, 2, 3, 4, 5, 6, 8, 9, 10}, but got %v", elements)
	}
}

func areSetEqual[T cmp.Ordered](xs, ys []T) bool {
	if len(xs) != len(ys) {
		return false
	}

	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	sort.Slice(ys, func(i, j int) bool { return ys[i] < ys[j] })

	for k, x := range xs {
		if ys[k] != x {
			return false
		}
	}

	return true
}
