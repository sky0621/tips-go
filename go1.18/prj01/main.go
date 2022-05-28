package main

import "fmt"

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

func main() {
	fmt.Println(addInts(1, 2))
	fmt.Println(addFloats(2.2, 3.3))
	fmt.Println(addString("a", "b"))

	fmt.Println(addAny[int](5, 6))
	fmt.Println(addAny[string]("7", "8"))

	o := &One{
		ID:   122,
		Name: "OneOne",
	}

	t := &Two{ID: 5656, IsAdmin: true}

	something(*o)
	something(*t)
}
