package main

import (
	"fmt"
)

func main() {
	r0 := restFunc()
	r0a := <-r0
	fmt.Println("r0a", r0a)

	r1a := <-r0
	fmt.Println("r1a", r1a)

	r2a := <-r0
	fmt.Println("r2a", r2a)

	r3a := <-r0
	fmt.Println("r3a", r3a)

	r4a := <-r0
	fmt.Println("r4a", r4a)
}

func restFunc() <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for i := 0; i < 3; i++ {
			result <- i
		}
	}()

	return result
}
