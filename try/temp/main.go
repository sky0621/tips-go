package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("<<< Listing 1 >>>")
	listing1()
	fmt.Println()
	fmt.Println("<<< Listing 2 >>>")
	listing2()
}

func listing1() {
	s := make([]int, 1)
	fmt.Println(s, len(s), cap(s))

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s1 := append(s, 1)
		fmt.Println(s1)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		s2 := append(s, 2)
		fmt.Println(s2)
	}(&wg)

	wg.Wait()
}

func listing2() {
	s := make([]int, 0, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s2 := append(sCopy, 2)
		fmt.Println(s2)
	}(&wg)

	wg.Wait()
}
