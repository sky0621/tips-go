package main

import "fmt"

func addAny[V int | float32 | string](a, b V) V {
	return a + b
}

func something[V Signal](v V) {
	fmt.Println(v)
	fmt.Println(v.doWhat("xxxx"))
}
