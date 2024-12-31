package main

import (
	"fmt"

	"github.com/sky0621/tips-go/how2use/library/standard/errgroup/client"

	svc "github.com/sky0621/tips-go/how2use/library/standard/errgroup"
)

func main() {
	if err := svc.NewService(client.NewDBClient()).Exec(); err != nil {
		fmt.Printf("failed to exec service: %+v", err)
	}
}
