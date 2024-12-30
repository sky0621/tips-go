package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan struct{}, 1)

	wg.Add(1)
	go sub("ping", ch, wg)

	wg.Add(1)
	go sub("     pong", ch, wg)

	// 最初のきっかけを与えてやる
	ch <- struct{}{}

	wg.Wait()
}

func sub(word string, ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		select {
		case <-ch:
			fmt.Println(word)
			time.Sleep(100 * time.Millisecond)
			ch <- struct{}{}
		}
	}
}
