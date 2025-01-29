package main

import (
	"fmt"
)

func main() {
	printInt(3)
	fmt.Println()

	mi := MyInt(8)
	printInt(mi)
	fmt.Println()

	mt := MyTypes[MyInt]{3, 5, 7}
	mt.doSomething()
}

func printInt[T ~int](x T) {
	fmt.Println(x)
}

type MyInt int

func (i *MyInt) Twice() {
	*i = *i * 2
}

type MyTypes[T MyInt] []T

func (m MyTypes[T]) doSomething() {
	for _, x := range m {
		fmt.Println(x)
	}
}
