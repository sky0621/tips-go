package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/functional/fp"
)

func main() {
	ary := []int{1, 2, 3}
	fmt.Printf("Remain([]int{1, 2, 3} -> %#v\n", fp.Remain(ary))
	ary2 := []int{4, 5, 6}
	fmt.Printf("Remain([]int{4, 5, 6} -> %#v\n", fp.Remain(ary2))
}
