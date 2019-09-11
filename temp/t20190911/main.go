package main

import (
	"fmt"
)

func main() {
	var n uint64
	for n = 1; n < 30; n++ {
		fmt.Printf("[%2d] %d\n", n, factorial(n, 1))
	}

	fmt.Println("")
	fmt.Println("===================================================================")
	fmt.Println("")

	for i := 1; i < 100; i++ {
		fizzbuzz(i)
	}

	fmt.Println("")
	fmt.Println("===================================================================")
	fmt.Println("")

	fmt.Println(fibo(20))
}

func factorial(n, sum uint64) uint64 {
	if n == 1 {
		return sum
	}
	return factorial(n-1, n*sum)
}

func fizzbuzz(n int) {
	switch {
	case n%15 == 0:
		fmt.Println("fizzbuzz")
	case n%3 == 0:
		fmt.Println("fizz")
	case n%5 == 0:
		fmt.Println("buzz")
	default:
		fmt.Println(n)
	}
}

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-2) + fibo(n-1)
}
