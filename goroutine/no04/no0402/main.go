package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go sub(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("END")
	}
}

func sub(ctx context.Context) {
	for {
		fmt.Printf("a ")
		time.Sleep(1 * time.Second)
	}
}
