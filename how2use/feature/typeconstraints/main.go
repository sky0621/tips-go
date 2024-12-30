package main

import (
	"fmt"
)

func main() {
	printInt(3)

	mi := MyInt(8)
	printInt(mi)

	mt := MyTypes[int]{3, 5, 7}
	mt.doSomething()
}

func printInt[T ~int](x T) {
	fmt.Println(x)
}

type MyInt int

type MyTypes[T any] []T

func (m MyTypes[T]) doSomething() {
	fmt.Println(m)
}
