package main

import (
	"fmt"
	"time"
)

// 非同期：非制御
// 同期のping実行終了後、非同期のpongが実行されるが、同時にメインスレッドが終了されるためpongの出力は無し
func main() {
	ping()
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
