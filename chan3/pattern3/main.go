package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		q <- "foo"
	}()
	go func() {
		time.Sleep(4 * time.Second)
		q <- "bar"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		q <- "mee"
	}()

	var cmds []string
	cmds = append(cmds, <-q) // とりあえず１つ受信

wait_some:
	for {
		select {
		case cmd := <-q: // 追加分を処理
			cmds = append(cmds, cmd)
		case <-time.After(1 * time.Second): // １秒超えたら処理しない！
			break wait_some
		}
	}

	for _, cmd := range cmds {
		fmt.Println(cmd)
	}
}
