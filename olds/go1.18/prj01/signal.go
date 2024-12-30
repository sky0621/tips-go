package main

type One struct {
	ID   int
	Name string
}

func (One) doWhat(s string) string {
	return s + "One"
}

type Two struct {
	ID      float32
	IsAdmin bool
}

func (Two) doWhat(s string) string {
	return s + "Two"
}

type Signal interface {
	One | Two

	doWhat(s string) string
}
