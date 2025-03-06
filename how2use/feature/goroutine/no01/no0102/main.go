package main

import (
	"fmt"
	"time"
)

// 非同期：非制御
// pingが非同期で、同期のpong実行中にpingも実行される
func main() {
	go ping()
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
