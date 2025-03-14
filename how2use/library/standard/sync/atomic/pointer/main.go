package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Data struct {
	value int
}

// go run -race ./main.go
func main() {
	/*
	 * race condition が発生しないだけであって、すべての数値が出力されることを保証するものではない。
	 */
	var p atomic.Pointer[Data]
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			p.Store(&Data{value: i})
			fmt.Println("Write:", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			data := p.Load()
			if data != nil {
				fmt.Println("Read:", data.value)
			}
		}()
	}

	wg.Wait()
}
