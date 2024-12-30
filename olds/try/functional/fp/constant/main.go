package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/functional/fp"
)

func main() {
	fmt.Printf("Constant(1) -> %#v\n", fp.Constant(1))
	fmt.Printf("Constant(999) -> %#v\n", fp.Constant(999))
	fmt.Printf("Constant(a) -> %#v\n", fp.Constant("a"))
}
