package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan string)

	disableCh := true
	//disableCh := false
	//ch1 <- 1

	var c chan int
	if disableCh {
		c = nil
	} else {
		c = ch1
	}

	// 一定時間経過後に ch2 にデータを送る
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	select {
	// c が nil の場合は、この case は実行されない
	case num := <-c:
		fmt.Println(num)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}
