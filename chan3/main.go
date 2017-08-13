package main

import (
	"fmt"
	"time"
)

func main() {
	task := make(chan string)
	taskquit := make(chan bool)
	workerquit := make(chan bool)

	go doTask(task, taskquit, workerquit)

	go callTask(task, taskquit)

	<-workerquit
	fmt.Println("Quit Worker!")
}

func doTask(t chan string, tq, wq chan bool) {
loop:
	for {
		select {
		case <-tq:
			fmt.Println("Quit Task")
			wq <- true
			break loop
		case job := <-t:
			fmt.Println(job)
		}
	}
}

func callTask(t chan string, tq chan bool) {
	for n := 0; n < 5; n++ {
		t <- fmt.Sprintf("Task %02d\n", n+1)
		time.Sleep(1 * time.Second)
	}
	tq <- true
}
