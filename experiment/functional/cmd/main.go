package main

import (
	"fmt"

	"github.com/sky0621/tips-go/experiment/functional"
)

func main() {
	// Constant
	fmt.Printf("Constant(1) -> %#v\n", functional.Constant(1))
	fmt.Printf("Constant(999) -> %#v\n", functional.Constant(999))
	fmt.Printf("Constant(a) -> %#v\n", functional.Constant("a"))

	// Empty
	var empty = functional.Empty
	fmt.Printf("empty(1) -> %#v\n", empty(1))
	fmt.Printf("empty(2) -> %#v\n", empty(2))
	fmt.Printf("empty(10) -> %#v\n", empty(10))
	fmt.Printf("empty(-5) -> %#v\n", empty(-5))

	// Add
	fmt.Printf("Add(3, 7) -> %d\n", functional.Add(3, 7))
	fmt.Printf("Add(18, 5) -> %d\n", functional.Add(18, 5))

	ary := []int{1, 2, 3}
	ary2 := []int{4, 5, 6}

	// First
	fmt.Printf("First([]int{1, 2, 3} -> %d\n", functional.First(ary))
	fmt.Printf("First([]int{4, 5, 6} -> %d\n", functional.First(ary2))

	// Remain
	fmt.Printf("Remain([]int{1, 2, 3} -> %#v\n", functional.Remain(ary))
	fmt.Printf("Remain([]int{4, 5, 6} -> %#v\n", functional.Remain(ary2))

	// Sum
	fmt.Printf("Sum([]int{1, 2, 3} -> %d\n", functional.Sum(ary))
}
