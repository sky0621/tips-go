package main

import (
	"fmt"
	"sync"
	"time"
)

// 同時処理実行可能数による非同期制御
func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan struct{}, 5)

	for i := 0; i < 25; i++ {
		wg.Add(1)

		ch <- struct{}{} // 10個までは詰められる（でも、それ以上はチャネルが空くまで詰められずにここで待機状態となる）
		go sub(i, ch, wg)
	}

	wg.Wait()
}

func sub(i int, ch chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		<-ch // 処理後にチャネルから値を抜き出さないと、次の goroutine が起動できない

		wg.Done()
	}()

	fmt.Println(i)
	time.Sleep(2 * time.Second)
}
