package main

import (
	"fmt"
	"time"
)

func main() {
	// バッファなしチャネル
	ch := make(chan int)

	pub := func(c chan int) {
		for {
			start := time.Now()
			c <- 1
			end := time.Now()
			fmt.Printf("[pub] %f秒\n", (end.Sub(start)).Seconds())
		}
	}

	sub := func(c chan int) {
		for {
			time.Sleep(2 * time.Second)
			start := time.Now()
			t := <-c
			fmt.Printf("          [sub] %d\n", t)
			end := time.Now()
			fmt.Printf("          [sub] %f秒\n", (end.Sub(start)).Seconds())
		}
	}

	go pub(ch)
	go sub(ch)

	time.Sleep(20 * time.Second)
}
