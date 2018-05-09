package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go exec(ch)

	no := 0
	for {
		time.Sleep(100 * time.Microsecond)
		if no > 100 {
			break
		}
		ch <- no
		no = no + 1
	}
}

func exec(c chan int) {
	for {
		fmt.Print(".")
		select {
		case s := <-c:
			fmt.Println("　")
			fmt.Println(s)
		default:
			// デフォルト節があると、ループの次の回へ。デフォルト節がない場合、ケース節のチャネルからの値取得待ち
		}
	}
}
