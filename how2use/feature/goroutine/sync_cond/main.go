package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mux := sync.Mutex{}
	cond := sync.NewCond(&mux)

	ready := false

	go func() {
		cond.L.Lock()
		defer cond.L.Unlock()
		for !ready {
			cond.Wait()
		}
		fmt.Println("Goroutine: Condition met!")
	}()

	time.Sleep(2 * time.Second)

	cond.L.Lock()
	ready = true
	cond.Signal() // 待機中のゴルーチン１つにシグナルを送る
	cond.L.Unlock()

	time.Sleep(1 * time.Second)
}
