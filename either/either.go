package types

type Either[A, B any] interface {
	IsLeft() bool
	IsRight() bool
	UnwrapLeft() (A, bool)
	UnwrapRight() (B, bool)
}

type either[A, B any] struct {
	left   A
	right  B
	isLeft bool
}

func Left[A, B any](value A) Either[A, B] {
	return &either[A, B]{left: value, isLeft: true}
}

func Right[A, B any](value B) Either[A, B] {
	return &either[A, B]{right: value, isLeft: false}
}

func (e *either[A, B]) IsLeft() bool {
	return e.isLeft
}

func (e *either[A, B]) IsRight() bool {
	return !e.isLeft
}

func (e *either[A, B]) UnwrapLeft() (A, bool) {
	return e.left, e.isLeft
}

func (e *either[A, B]) UnwrapRight() (B, bool) {
	return e.right, !e.isLeft
}
