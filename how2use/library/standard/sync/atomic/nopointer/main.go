package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	value int
}

// go run -race ./main.go
func main() {
	var data *Data
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data = &Data{value: i}
			fmt.Println("Write:", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			if data != nil {
				fmt.Println("Read:", data.value)
			}
		}()
	}

	wg.Wait()
}
