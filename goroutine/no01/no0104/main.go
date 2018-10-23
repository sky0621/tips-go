package main

import (
	"fmt"
	"time"
)

func main() {
	go ping()
	go pong()
}

func ping() {
	for i := 0; i < 5; i++ {
		fmt.Println("ping")
		time.Sleep(100 * time.Millisecond)
	}
}

func pong() {
	for i := 0; i < 5; i++ {
		fmt.Println("     pong")
		time.Sleep(100 * time.Millisecond)
	}
}
