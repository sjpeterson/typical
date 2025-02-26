package set

import (
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

func TestSet_IsEqualTo(t *testing.T) {
	testCases := []struct {
		first         Set[int]
		second        Set[int]
		expectedEqual bool
	}{
		{NewSet[int](), NewSet[int](), true},
		{NewSet(3), NewSet(3), true},
		{NewSet(5, 7, 2), NewSet(2, 7, 5, 5, 2), true},
		{NewSet[int](), NewSet(5, 6, 2), false},
		{NewSet(5, 6, 2), NewSet[int](), false},
		{NewSet(6, 3, 1, 4), NewSet(6, 3, 1, 1, 8, 4), false},
	}

	for _, testCase := range testCases {
		if equal := testCase.first.IsEqualTo(testCase.second); equal != testCase.expectedEqual {
			t.Errorf("expected %v and %v equality to be %v, but got %v", testCase.first, testCase.second, testCase.expectedEqual, equal)
		}
	}
}

func TestSet_IsEmpty(t *testing.T) {
	emptySet := NewSet[int]()
	nonEmptySet := NewSet(1, 5, 9)

	if !emptySet.IsEmpty() {
		t.Error("expected IsEmpty to return true for empty set, but got false")
	}

	if nonEmptySet.IsEmpty() {
		t.Error("expected IsEmpty to return false for non-empty set, but got true")
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

func TestSet_Clone(t *testing.T) {
	testSet := NewSet(6, 8, 3, 5, 2)
	clonedSet := testSet.Clone()

	if !clonedSet.IsEqualTo(testSet) {
		t.Errorf("expected cloned set to be equal to the original")
	}

	clonedSet.Add(1)
	if clonedSet.IsEqualTo(testSet) {
		t.Errorf("cloned set still equal to original after adding an element")
	}
}

func TestSet_Elements(t *testing.T) {
	emptySet := NewSet[int]()
	singletonSet := NewSet(5)
	primesBelowFifteen := NewSet(2, 3, 5, 7, 11, 13)

	if elements := emptySet.Elements(); !areSetEqual(elements, []int{}) {
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

func TestSet_Discard(t *testing.T) {
	a := NewSet(2, 4, 7, 8)

	a.Discard(7)

	if elements := a.Elements(); !areSetEqual(elements, []int{2, 4, 8}) {
		t.Errorf("expected a to be {2, 4, 8}, but got %v", elements)
	}
}

func TestUnion(t *testing.T) {
	testCases := []struct {
		sets          []Set[int]
		expectedUnion Set[int]
	}{
		{[]Set[int]{}, NewSet[int]()},
		{[]Set[int]{NewSet(2, 4, 6, 8)}, NewSet(2, 4, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5)}, NewSet(1, 2, 3, 4, 5, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5), NewSet(2, 9, 10)}, NewSet(1, 2, 3, 4, 5, 6, 8, 9, 10)},
	}
	for _, testCase := range testCases {
		if union := Union(testCase.sets...); !union.IsEqualTo(testCase.expectedUnion) {
			t.Errorf("expected the union of %v to be %v, but got %v", testCase.sets, testCase.expectedUnion, union)
		}
	}
}

func TestIntersection(t *testing.T) {
	testCases := []struct {
		sets                 []Set[int]
		expectedIntersection Set[int]
	}{
		{[]Set[int]{}, NewSet[int]()},
		{[]Set[int]{NewSet(2, 4, 6, 8)}, NewSet(2, 4, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5)}, NewSet(2, 4)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5), NewSet(2, 9, 10)}, NewSet(2)},
	}
	for _, testCase := range testCases {
		if intersection := Intersection(testCase.sets...); !intersection.IsEqualTo(testCase.expectedIntersection) {
			t.Errorf("expected the intersection of %v to be %v, but got %v", testCase.sets, testCase.expectedIntersection, intersection)
		}
	}
}

func TestDifference(t *testing.T) {
	testCases := []struct {
		sets               []Set[int]
		expectedDifference Set[int]
	}{
		{[]Set[int]{}, NewSet[int]()},
		{[]Set[int]{NewSet(2, 4, 6, 8)}, NewSet(2, 4, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5)}, NewSet(6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5), NewSet(2, 6, 10)}, NewSet(8)},
	}
	for _, testCase := range testCases {
		if difference := Difference(testCase.sets...); !difference.IsEqualTo(testCase.expectedDifference) {
			t.Errorf("expected the difference of %v to be %v, but got %v", testCase.sets, testCase.expectedDifference, difference)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	testCases := []struct {
		sets               []Set[int]
		expectedDifference Set[int]
	}{
		{[]Set[int]{}, NewSet[int]()},
		{[]Set[int]{NewSet(2, 4, 6, 8)}, NewSet(2, 4, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5)}, NewSet(1, 3, 5, 6, 8)},
		{[]Set[int]{NewSet(2, 4, 6, 8), NewSet(1, 2, 3, 4, 5), NewSet(3, 6, 10)}, NewSet(1, 5, 8, 10)},
	}
	for _, testCase := range testCases {
		if symmetricDifference := SymmetricDifference(testCase.sets...); !symmetricDifference.IsEqualTo(testCase.expectedDifference) {
			t.Errorf("expected the symmetric difference of %v to be %v, but got %v", testCase.sets, testCase.expectedDifference, symmetricDifference)
		}
	}
}

func areSetEqual(xs, ys []int) bool {
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
