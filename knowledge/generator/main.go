package main

import (
	"fmt"
)

func main() {
	generator := fiveValueGenerator()
	for i := 0; i < 8; i++ {
		v := <-generator
		fmt.Println(v)
	}
}

func fiveValueGenerator() <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := 1; i < 6; i++ {
			result <- i
		}
	}()
	return result
}
