package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, &wg)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		default:
			//
		}
		fmt.Println("goroutine: 処理")
		time.Sleep(1 * time.Second)
	}
}