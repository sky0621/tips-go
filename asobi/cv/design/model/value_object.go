package model

type ValueObject[T any] interface {
	GetValue() T
}
