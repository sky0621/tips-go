package main

import "fmt"

type If interface {
	exec()
}

type Str1 struct {
}

func (s Str1) exec() {
	fmt.Println("I'm Str1")
}

type Str2 struct {
}

func (s Str2) exec() {
	fmt.Println("I'm Str2")
}

func ifExec[T If](ts []T) {
	for _, t := range ts {
		t.exec()
	}
}

func main() {
	ifExec([]If{Str1{}, Str2{}})
}
