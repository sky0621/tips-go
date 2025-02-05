package main

import "fmt"

func main() {
	var a A
	a.test()
}

type A int

func (a A) test() {
	fmt.Println("a:", a)
}
