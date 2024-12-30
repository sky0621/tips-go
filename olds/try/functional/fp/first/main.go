package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/functional/fp"
)

func main() {
	ary := []int{1, 2, 3}
	fmt.Printf("First([]int{1, 2, 3} -> %d\n", fp.First(ary))
	ary2 := []int{4, 5, 6}
	fmt.Printf("First([]int{4, 5, 6} -> %d\n", fp.First(ary2))
}
