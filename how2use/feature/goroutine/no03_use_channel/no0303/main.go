package main

import (
	"fmt"
	"sync"
	"time"
)

// 5個のwordがランダムに1つのチャネルを奪い合い出力する
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
		go sub(word, ch, wg)
	}

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
			time.Sleep(1 * time.Second)
			ch <- struct{}{}
		}
	}
}
