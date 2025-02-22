package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go exec1(ch)
	go exec2(ch2)

	for {
		select {
		case i, ok := <-ch:
			if ok {
				fmt.Printf("[CH] %d\n", i)
				continue
			}
			ch = nil
		case j, ok := <-ch2:
			if ok {
				fmt.Printf("[CH2] %d\n", j)
				continue
			}
			ch2 = nil
		}
		if ch == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("main end")
}

func exec1(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		if i%2 == 0 {
			ch <- i
		}
	}
	close(ch)
	fmt.Println("fin exec1")
}

func exec2(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(150 * time.Millisecond)
		if i%2 != 0 {
			ch <- i
		}
	}
	close(ch)
	fmt.Println("fin exec2")
}
