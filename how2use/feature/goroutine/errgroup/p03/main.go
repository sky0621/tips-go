package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)

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

	for _, task := range tasks {
		group.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return task()
			}
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println("All tasks completed successfully.")
	}
}
