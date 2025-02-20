package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		s := NewRectangle(i, 2*i)
		fmt.Println(s.Area())
	}
}

func NewRectangle(width, height int) Rectangle {
	return Rectangle{width, height}
}

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}
