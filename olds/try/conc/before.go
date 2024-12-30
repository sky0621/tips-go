package main

import (
	"fmt"
	"sync"
	"time"
)

func before() {
	fmt.Println("Start")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go doSomethingThatMightPanic(&wg)
	//go func() {
	//	defer func() {
	//		if v := recover(); v != nil {
	//			fmt.Println("panic:", v)
	//		}
	//wg.Done()
	//}()
	//doSomethingThatMightPanic(&wg)
	//doSomethingThatMightPanic()
	//}()
	wg.Wait()
	fmt.Println("End")
}

// func doSomethingThatMightPanic() {
func doSomethingThatMightPanic(wg *sync.WaitGroup) {
	fmt.Println("start doSomethingThatMightPanic")
	defer func() {
		fmt.Println("waitGroup.Done")
		wg.Done()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("before panic")
	panic("panic")
}
