package main

import (
	"error-resolver/functional/fp"
	"fmt"
)

func main() {
	res := fp.Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
}
