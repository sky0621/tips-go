package main

import (
	"fmt"
	"time"
)

// 同期
// 順番通りに実行される
func main() {
	ping()
	pong()
}

func ping() {
	for i := 0; i < 5; i++ {
		fmt.Println("ping")
		time.Sleep(1 * time.Second)
	}
}

func pong() {
	for i := 0; i < 5; i++ {
		fmt.Println("     pong")
		time.Sleep(1 * time.Second)
	}
}
