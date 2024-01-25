package set

type empty struct{}

// Set is a unordered collection of unique elements
type Set[T comparable] interface {
	Add(element T)
	Cardinality() int
	Contains(element T) bool
	DoesNotContain(element T) bool
	Elements() []T
	Union(other Set[T])
	Intersection(other Set[T])
	Difference(other Set[T])
	SymmetricDifference(other Set[T])
}

type set[T comparable] struct {
	elements map[T]empty
}

// New creates a new set with the provided elements
func New[T comparable](elements ...T) Set[T] {
	setElements := make(map[T]empty)
	for _, element := range elements {
		setElements[element] = empty{}
	}

	return &set[T]{elements: setElements}
}

func (s *set[T]) Add(element T) {
	s.elements[element] = empty{}
}

// Cardinality is the number of elements in the set
func (s *set[T]) Cardinality() int {
	return len(s.elements)
}

// Contains returns true if the element belongs to the set
func (s *set[T]) Contains(element T) bool {
	_, ok := s.elements[element]

	return ok
}

// DoesNotContain returns true if the element does not belong to the set
func (s *set[T]) DoesNotContain(element T) bool {
	_, ok := s.elements[element]

	return !ok
}

// Elements returns the elements in the set in a slice
func (s *set[T]) Elements() []T {
	elements := make([]T, 0, len(s.elements))
	for element, _ := range s.elements {
		elements = append(elements, element)
	}

	return elements
}

// Union updates the set to be the union of itself and another set
func (s *set[T]) Union(other Set[T]) {
	for _, element := range other.Elements() {
		s.elements[element] = empty{}
	}
}

// Intersection updates the set to be the intersection of itself and another set
func (s *set[T]) Intersection(other Set[T]) {
	for element, _ := range s.elements {
		if other.DoesNotContain(element) {
			delete(s.elements, element)
		}
	}
}

// Difference updates the set to be the set difference of itself and another set
func (s *set[T]) Difference(other Set[T]) {
	for _, element := range other.Elements() {
		delete(s.elements, element)
	}
}

// SymmetricDifference updates the set to be the symmetric difference of itself and another set
func (s *set[T]) SymmetricDifference(other Set[T]) {
	for _, element := range other.Elements() {
		if _, ok := s.elements[element]; ok {
			delete(s.elements, element)
		} else {
			s.elements[element] = empty{}
		}
	}
}

// Union computes the union of zero or more sets
func Union[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return New[T]()
	}
	set := New(sets[0].Elements()...)

	for _, other := range sets[1:] {
		set.Union(other)
	}

	return set
}
