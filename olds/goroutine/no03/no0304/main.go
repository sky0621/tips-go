package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan struct{}, 1)

	words := []string{
		"ping",
		"     pong",
		"          pang",
		"               peng",
		"                    pung",
	}

	for _, word := range words {
		wg.Add(1)
		go sub(word, ch, wg)
	}

	// 最初のきっかけを与えてやる
	ch <- struct{}{}

	wg.Wait()
}

func sub(word string, ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(word)
		time.Sleep(100 * time.Millisecond)
	}
}
