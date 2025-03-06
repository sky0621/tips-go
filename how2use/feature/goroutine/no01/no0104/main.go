package main

import (
	"fmt"
	"time"
)

// 非同期：非制御
// pingもpongも実行開始されて同時にメインスレッドが終了されるため何も出力されない
func main() {
	go ping()
	go pong()
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
