package main

import (
	"fmt"
)

func main() {
	var n uint64
	for n = 1; n < 30; n++ {
		fmt.Printf("[%2d] %d\n", n, factorial(n, 1))
	}
	var o uint64 = 14197454024290336768
	fmt.Println(o * 22)
}

func factorial(n, sum uint64) uint64 {
	if n == 1 {
		return sum
	}
	return factorial(n-1, n*sum)
}
