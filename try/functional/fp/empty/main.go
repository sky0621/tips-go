package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/functional/fp"
)

func main() {
	var empty = fp.Empty
	fmt.Printf("empty(1) -> %#v\n", empty(1))
	fmt.Printf("empty(2) -> %#v\n", empty(2))
	fmt.Printf("empty(10) -> %#v\n", empty(10))
	fmt.Printf("empty(-5) -> %#v\n", empty(-5))
}
