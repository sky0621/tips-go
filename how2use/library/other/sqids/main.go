package main

import (
	"fmt"
	"github.com/sqids/sqids-go"
)

func main() {
	s, _ := sqids.New()
	id, _ := s.Encode([]uint64{77, 99, 111})
	fmt.Println(id)
	decoded := s.Decode(id)
	fmt.Println(decoded)
}
