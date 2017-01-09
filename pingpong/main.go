package main

import "fmt"

func main() {
	pingc := make(chan int)
	pongc := make(chan int)
	quitc := make(chan int)
	go pong(pingc, pongc, quitc)
	go ping(3, pingc, pongc)
	<- quitc
}

func ping(n int, ping, pong chan int) {
	for {
		pong <- n
		if n == 0 {
			fmt.Println("ping finished")
			return
		}
		<- pong
		n = n + 1
	}
}

func pong(ping, pong, quit chan int) {
	for {
		n := <- pong
		if n == 0 {
			fmt.Println("pong finished")
			quit <- 0
			return
		}
		fmt.Println("pong sending ", n)
		ping <- n
	}
}

