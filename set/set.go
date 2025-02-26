package set

type empty struct{}

// Set is a unordered collection of unique elements.
type Set[T comparable] interface {
	Add(element T)
	Discard(element T)
	Clone() Set[T]
	Cardinality() int
	Contains(element T) bool
	DoesNotContain(element T) bool
	IsEmpty() bool
	Elements() []T
	Union(other Set[T])
	Intersection(other Set[T])
	Difference(other Set[T])
	SymmetricDifference(other Set[T])
	IsEqualTo(other Set[T]) bool
}

type set[T comparable] struct {
	elements map[T]empty
}

// NewSet creates a new set with the provided elements.
func NewSet[T comparable](elements ...T) Set[T] {
	setElements := make(map[T]empty)
	for _, element := range elements {
		setElements[element] = empty{}
	}

	return &set[T]{elements: setElements}
}

// Clone creates a clone of the set
func (s *set[T]) Clone() Set[T] {
	set := NewSet[T]()
	for element := range s.elements {
		set.Add(element)
	}

	return set
}

// Add adds an element to the set.
func (s *set[T]) Add(element T) {
	s.elements[element] = empty{}
}

// Discard removes an element from the set if it is a member. If it is not a member, do nothing.
func (s *set[T]) Discard(element T) {
	delete(s.elements, element)
}


// Cardinality is the number of elements in the set.
func (s *set[T]) Cardinality() int {
	return len(s.elements)
}

// Contains returns true if the element belongs to the set.
func (s *set[T]) Contains(element T) bool {
	_, ok := s.elements[element]

	return ok
}

// DoesNotContain returns true if the element does not belong to the set.
func (s *set[T]) DoesNotContain(element T) bool {
	_, ok := s.elements[element]

	return !ok
}

// Elements returns the elements in the set in a slice.
func (s *set[T]) Elements() []T {
	elements := make([]T, 0, len(s.elements))
	for element := range s.elements {
		elements = append(elements, element)
	}

	return elements
}

// Union updates the set to be the union of itself and another set.
func (s *set[T]) Union(other Set[T]) {
	for _, element := range other.Elements() {
		s.elements[element] = empty{}
	}
}

// Intersection updates the set to be the intersection of itself and another set.
func (s *set[T]) Intersection(other Set[T]) {
	for element := range s.elements {
		if other.DoesNotContain(element) {
			delete(s.elements, element)
		}
	}
}

// Difference updates the set to be the set difference of itself and another set.
func (s *set[T]) Difference(other Set[T]) {
	for _, element := range other.Elements() {
		delete(s.elements, element)
	}
}

// SymmetricDifference updates the set to be the symmetric difference of itself and another set.
func (s *set[T]) SymmetricDifference(other Set[T]) {
	for _, element := range other.Elements() {
		if _, ok := s.elements[element]; ok {
			delete(s.elements, element)
		} else {
			s.elements[element] = empty{}
		}
	}
}

// IsEmpty returns true if the set is the empty set.
func (s *set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// IsEqualTo returns true if the set is equal to another set.
func (s *set[T]) IsEqualTo(other Set[T]) bool {
	if s.Cardinality() != other.Cardinality() {
		return false
	}
	for _, element := range other.Elements() {
		if _, ok := s.elements[element]; !ok {
			return false
		}
	}
	return true
}

// Union computes the union of zero or more sets.
func Union[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	set := NewSet(sets[0].Elements()...)

	for _, other := range sets[1:] {
		set.Union(other)
	}

	return set
}

// Intersection computes the intersection of zero or more sets.
func Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	set := NewSet(sets[0].Elements()...)

	for _, other := range sets[1:] {
		set.Intersection(other)
	}

	return set
}

// Difference computes the set difference of zero or more sets.
// When more than two sets are provided, the set of elements that are in the
// first set but in neither of the subsequent sets is returned.
func Difference[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}
	set := NewSet(sets[0].Elements()...)
	if len(sets) > 1 {
		set.Difference(Union(sets[1:]...))
	}

	return set
}

// SymmetricDifference computes the symmetric difference of zero or more sets.
// When more than two sets are provided, the set of elements that are in
// exactly one of the sets is returned.
func SymmetricDifference[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return NewSet[T]()
	}

	unionOfAll := Union(sets...)
	unionOfIntersections := NewSet[T]()
	nSets := len(sets)
	for m := 0; m < nSets; m++ {
		for n := m + 1; n < nSets; n++ {
			unionOfIntersections.Union(Intersection(sets[m], sets[n]))
		}
	}

	return Difference(unionOfAll, unionOfIntersections)
}
