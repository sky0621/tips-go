package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		q <- 1
	}()

	<-q
	fmt.Println("End")
}
