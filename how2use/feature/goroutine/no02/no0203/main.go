package main

import (
	"fmt"
	"time"
)

// チャネルの制御によりpingもpongも終了してからメインスレッドも終了する
func main() {
	ch := make(chan struct{}, 2)
	go ping(ch)
	go pong(ch)
	<-ch
}

func ping(ch chan struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println("ping")
		time.Sleep(100 * time.Millisecond)
	}
	ch <- struct{}{}
}

func pong(ch chan struct{}) {
	for i := 0; i < 5; i++ {
		fmt.Println("     pong")
		time.Sleep(100 * time.Millisecond)
	}
	ch <- struct{}{}
}
