package model

type ValueObject[V any] interface {
	Value() V
}
