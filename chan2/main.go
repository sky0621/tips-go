package main

import (
	"fmt"
	"time"
)

func main() {
	no := 0
	ch := make(chan int)
	fmt.Println("channel [ch] made")
	end := make(chan bool)
	fmt.Println("channel [end] made")
loopend:
	for {
		time.After(3 * time.Second)
		no = no + 1
		fmt.Printf("no [%d]\n", no)
		go meth(no, ch, end)
		fmt.Println("after go meth")
		time.After(3 * time.Second)
		ch <- no

		<-end
		goto loopend
	}
	// fmt.Println("All Finished")
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
