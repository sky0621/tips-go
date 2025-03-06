package main

import (
	"fmt"
	"time"
)

// 3秒スリープの間にpingとpongが出力される
func main() {
	go ping()
	go pong()
	time.Sleep(6 * time.Second)
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
