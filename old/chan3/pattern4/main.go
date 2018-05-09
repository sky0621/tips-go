package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	task := make(chan string, 5) // ５個のタスク

	// ３つのワーカー
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(&wg, task)
	}

	task <- "t01"
	task <- "t02"
	task <- "t03"
	task <- "t04"
	task <- "t05"
	task <- "t06"
	task <- "t07"
	close(task)

	wg.Wait()
}

func work(wg *sync.WaitGroup, task chan string) {
	defer wg.Done()
	for {
		t, ok := <-task
		if !ok { // チャネルがクローズされたら ok は false
			return
		}
		fmt.Println(t)
		time.Sleep(2 * time.Second)
	}
}
