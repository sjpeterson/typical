package maybe

// Maybe[T] is a value of type T or no value.
type Maybe[T any] interface {
	IsSome() bool
	IsNone() bool
	Unwrap() (T, bool)
}

type maybe[T any] struct {
	value  T
	isSome bool
}

// Some creates a Maybe with a value.
func Some[T any](value T) Maybe[T] {
	return &maybe[T]{value, true}
}

// None creates a Maybe without a value.
func None[T any]() Maybe[T] {
	var zero T

	return &maybe[T]{zero, false}
}

// IsSome returns true if the Maybe has a valid value.
func (m *maybe[T]) IsSome() bool {
	return m.isSome
}

// IsNone returns true if the Maybe does not have a valid value.
func (m *maybe[T]) IsNone() bool {
	return !m.isSome
}

// Unwrap returns the value (if any) and a bool indicating whether or not it is valid.
func (m *maybe[T]) Unwrap() (T, bool) {
	return m.value, m.isSome
}
