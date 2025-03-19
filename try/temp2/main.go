package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	tac := time.After(2 * time.Second)

	t := <-tac
	fmt.Println(t)
}
