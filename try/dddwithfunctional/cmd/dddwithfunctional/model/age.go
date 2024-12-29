package model

import "errors"

var _ ValueObject[int] = (*Age)(nil)

type Age struct {
	value int
}

func (a *Age) Value() int {
	return a.value
}

func NewAge(value int) (*Age, error) {
	if value < 0 {
		return nil, errors.New("age cannot be negative")
	}
	if value > 150 {
		return nil, errors.New("age is too large")
	}
	return &Age{value: value}, nil
}
