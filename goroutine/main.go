package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 10

const maxProcesses = 3

func main() {

	sema := make(chan int, maxProcesses)

	notify := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(no int, sema chan int, notify chan<- int) {
			sema <- 0
			time.Sleep(time.Duration(rand.Int63n(3)) * time.Second)
			fmt.Println("処理完了", no)
			<-sema
			notify <- no
		}(i, sema, notify)
	}

	for i := 0; i < goroutines; i++ {
		<-notify
	}

	fmt.Println("すべて完了")
}
