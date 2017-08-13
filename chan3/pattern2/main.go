package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int, 2)

	go func() {
		time.Sleep(3 * time.Second)
		q <- 1
	}()

	for {
		if len(q) > 0 {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Something!")
	}
	// <-q
	fmt.Println("End")
}
