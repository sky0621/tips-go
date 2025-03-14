package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroupの制御によりpingもpongも終了してからメインスレッドも終了する
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go ping(wg)
	wg.Add(1)
	go pong(wg)
	wg.Wait()
}

func ping(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("ping")
		time.Sleep(1 * time.Second)
	}
}

func pong(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("     pong")
		time.Sleep(1 * time.Second)
	}
}
