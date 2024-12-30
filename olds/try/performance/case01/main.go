package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("test")
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}
	for _, c := range s {
		go func(c string) {
			for {
				fmt.Print(c)
				time.Sleep(500 * time.Millisecond)
			}
		}(c)
	}
	time.Sleep(30 * time.Second)
}
