package either

type Either[L any, R any] struct {
	left  *L
	right *R
}

func Left[L any, R any](l L) Either[L, R] {
	return Either[L, R]{left: &l}
}

func Right[L any, R any](r R) Either[L, R] {
	return Either[L, R]{right: &r}
}

func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

func (e Either[L, R]) LeftValue() (L, bool) {
	if e.left != nil {
		return *e.left, true
	}
	var zero L
	return zero, false
}

func (e Either[L, R]) RightValue() (R, bool) {
	if e.right != nil {
		return *e.right, true
	}
	var zero R
	return zero, false
}
