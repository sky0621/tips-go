package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	results, err := handle(context.Background(), []Circle{{1}, {2}, {3}, {4}, {5}})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func handle(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	group, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		group.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

type (
	Circle struct{ Num int }
	Result struct{ Num int }
)

func foo(ctx context.Context, c Circle) (Result, error) {
	if c.Num%2 == 0 {
		return Result{}, fmt.Errorf("error: %d", c.Num)
	}
	return Result{Num: c.Num}, nil
}
