package main

import (
	"fmt"

	svc "github.com/sky0621/tips-go/try/errgroup"

	"github.com/sky0621/tips-go/try/errgroup/client"
)

func main() {
	if err := svc.NewService(client.NewDBClient()).Exec(); err != nil {
		fmt.Printf("failed to exec service: %+v", err)
	}
}
