package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()

	tasks := []func() error{
		func() error {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("task 1 finished")
			return nil
		},
		func() error {
			return fmt.Errorf("task e1 encountered an error")
		},
		func() error {
			time.Sleep(800 * time.Millisecond)
			fmt.Println("task 2 finished")
			return nil
		},
		func() error {
			return fmt.Errorf("task e2 encountered an error")
		},
		func() error {
			time.Sleep(1200 * time.Millisecond)
			fmt.Println("task 3 finished")
			return nil
		},
		func() error {
			return fmt.Errorf("task e3 encountered an error")
		},
	}

	wg := sync.WaitGroup{}
	errs := make([]error, len(tasks))
	for _, task := range tasks {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				errs = append(errs, ctx.Err())
			default:
				errs = append(errs, task())
			}
		}(&wg)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			break
		}
	}
}
