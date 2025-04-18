package main

import (
	"fmt"
	"github.com/sqids/sqids-go"
)

func main() {
	s, _ := sqids.New()
	id, _ := s.Encode([]uint64{77, 99, 111})
	// W1x0jPr1k
	fmt.Println(id)
	decoded := s.Decode(id)
	fmt.Println(decoded)

	id2, _ := s.Encode([]uint64{1})
	// Uk
	fmt.Println(id2)

	s2, _ := sqids.New(sqids.Options{MinLength: 8})
	id3, _ := s2.Encode([]uint64{1})
	// UkLWZg9D
	fmt.Println(id3)
}
