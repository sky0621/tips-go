package main

import (
	"fmt"
	"sync"
	"time"
)

// チャネルで制御しないのでランダムに5個のwordが出力される
func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan struct{}, 1)

	words := []string{
		"A",
		" B",
		"  C",
		"   D",
		"    E",
	}

	for _, word := range words {
		wg.Add(1)
		go sub(word, wg)
	}

	// 最初のきっかけを与えてやる
	ch <- struct{}{}

	wg.Wait()
}

func sub(word string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(word)
		time.Sleep(100 * time.Millisecond)
	}
}
