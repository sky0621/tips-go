package main

import (
	"fmt"
	"time"
)

func main() {
	no := 0
	ch := make(chan int)
	end := make(chan bool)
loopend:
	for {
		time.After(3 * time.Second)
		no = no + 1
		go meth(no, ch, end)
		time.After(3 * time.Second)
		ch <- no

		<-end
		goto loopend
	}
}

func meth(no int, c chan int, end chan bool) {
	fmt.Println(no)
	res := <-c
	fmt.Println(res)
	if no > 50 {
		end <- true
	}
	close(c)
}
