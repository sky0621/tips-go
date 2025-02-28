package main

import (
	"sync"

	util "github.com/sky0621/tips-go/knowledge/goroutine_throttling"
)

func main() {
	util.ExecMain(func() {

		limit := make(chan struct{}, 5)

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			i := i
			go func() {
				limit <- struct{}{} // 20個詰め終わったあとはブロック
				defer wg.Done()
				util.Download(i)
				<-limit
			}()
		}
		wg.Wait()
	})
}
