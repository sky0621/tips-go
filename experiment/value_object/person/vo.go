package person

import (
	"fmt"
	"reflect"
)

type Validate[T any] func(v T) bool

type ValueObject[T any] interface {
	Value() T
	Equals(other ValueObject[T]) bool
	String() string
}

func NewValueObject[T any](value T, validators ...Validate[T]) (ValueObject[T], error) {
	for _, validate := range validators {
		if !validate(value) {
			return nil, fmt.Errorf("value object validation failed")
		}
	}
	return &valueObject[T]{value: value}, nil
}

type valueObject[T any] struct {
	value T
}

func (v valueObject[T]) Value() T {
	return v.value
}

func (v valueObject[T]) Equals(other ValueObject[T]) bool {
	return reflect.DeepEqual(v.value, other.Value())
}

func (v valueObject[T]) String() string {
	return fmt.Sprintf("value: %v", v.value)
}
