package tuples

type Pair[A any, B any] struct {
	First  A
	Second B
}

// NewPair is a convenience function to create a new Pair without instantiation.
func NewPair[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}

// Destructure provides ergonomic destructuring of a Pair without needing to name the fields.
func (p Pair[A, B]) Destructure() (A, B) {
	return p.First, p.Second
}

type Tup3[A any, B any, C any] struct {
	First  A
	Second B
	Third  C
}

func NewTup3[A, B, C any](first A, second B, third C) Tup3[A, B, C] {
	return Tup3[A, B, C]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

// Destructure provides ergonomic destructuring of a Tup3 without needing to name the fields.
func (t Tup3[A, B, C]) Destructure() (A, B, C) {
	return t.First, t.Second, t.Third
}

type Tup4[A any, B any, C any, D any] struct {
	First  A
	Second B
	Third  C
	Fourth D
}

// NewTup4 is a convenience function to create a new Pair without instantiation.
func NewTup4[A, B, C, D any](first A, second B, third C, fourth D) Tup4[A, B, C, D] {
	return Tup4[A, B, C, D]{
		First:  first,
		Second: second,
		Third:  third,
		Fourth: fourth,
	}
}

// Destructure provides ergonomic destructuring of a Tup4 without needing to name the fields.
func (t Tup4[A, B, C, D]) Destructure() (A, B, C, D) {
	return t.First, t.Second, t.Third, t.Fourth
}
