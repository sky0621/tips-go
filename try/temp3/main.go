package main

import (
	"errors"
	"fmt"
	"github.com/sky0621/tips-go/try/temp3/a"
)

func main() {
	r := a.Aa()
	fmt.Println(r)
	r2 := errors.Unwrap(r)
	fmt.Println(r2)
}
