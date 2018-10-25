package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// メインでの待機には WaitGroup を使う
	wg := &sync.WaitGroup{}

	// あえて明示的に「1」を指定
	pingCh := make(chan struct{}, 1)
	pongCh := make(chan struct{}, 1)

	wg.Add(1)
	go ping(pingCh, pongCh, wg)
	wg.Add(1)
	go pong(pingCh, pongCh, wg)

	// 最初のきっかけを与えてやる
	pingCh <- struct{}{}

	wg.Wait()
}

func ping(pingCh chan struct{}, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("ping")
		time.Sleep(100 * time.Millisecond)

		select {
		case <-pingCh:
			pongCh <- struct{}{}
		}
	}
	fmt.Println("ping fin")
}

func pong(pingCh chan struct{}, pongCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("     pong")
		time.Sleep(100 * time.Millisecond)
		select {
		case <-pongCh:
			pingCh <- struct{}{}
		}
	}
	fmt.Println("pong fin")
}
