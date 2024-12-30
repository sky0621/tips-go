package main

type Vector[T any] []T

func (v *Vector[T]) Push(x T) { *v = append(*v, x) }
