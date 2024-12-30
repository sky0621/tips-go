package main

import (
	"fmt"
	"tips-go/try/stringer/crud"
)

func main() {
	c := &C{}
	c.Show()

	fmt.Println(Red)
	fmt.Println(Yellow)
	fmt.Println(Green)
	fmt.Println(Red == Green)
	fmt.Println(Red == 3)
	var three int = 3
	var threeSignal Signal = Signal(three)
	fmt.Println(Red == threeSignal)

	fmt.Println()

	fmt.Println(crud.Create == crud.Read)
	fmt.Println(crud.Update == crud.Update)
	fmt.Println(crud.Delete)
	fmt.Println(crud.Other)
}

type A interface {
	Show()
}
type B interface {
	Show()
}

type C struct {
}

func (c *C) Show() {
	fmt.Println("c.Show")
}
