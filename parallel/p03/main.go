package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	// 終了通知用のコンテキスト生成
	ctx, cancel := context.WithCancel(context.Background())

	go func(fn context.CancelFunc) {
		time.Sleep(3 * time.Second)
		fmt.Println("do cancel()")
		fn()
	}(cancel)

	// キャンセル通知を待機
	fmt.Println("before cancel call")
	<-ctx.Done()
	fmt.Println("after  cancel call")

	fmt.Println("End")
}
