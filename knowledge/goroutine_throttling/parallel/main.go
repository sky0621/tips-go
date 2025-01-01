package main

import (
	"sync"

	util "github.com/sky0621/tips-go/knowledge/goroutine_throttling"
)

func main() {
	util.ExecMain(func() {
		var wg sync.WaitGroup
		for i := 0; i < 50; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				util.Download(i)
			}()
		}
		wg.Wait()
	})
}
