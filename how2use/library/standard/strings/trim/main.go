package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "123oxo"
	fmt.Println(s)

	fmt.Printf("strings.Trim(s, \"xo\"): %s\n", strings.Trim(s, "xo"))
	fmt.Printf("strings.TrimLeft(s, \"xo\"): %s\n", strings.TrimLeft(s, "xo"))
	fmt.Printf("strings.TrimRight(s, \"xo\"): %s\n", strings.TrimRight(s, "xo"))
	fmt.Printf("strings.TrimPrefix(s, \"xo\"): %s\n", strings.TrimPrefix(s, "xo"))
	fmt.Printf("strings.TrimSuffix(s, \"xo\"): %s\n", strings.TrimSuffix(s, "xo"))
}
