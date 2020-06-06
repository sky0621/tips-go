package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/errgroup/client"
)

func main() {
	oc := client.NewOrderClient()
	mc := client.NewMailClient()
	s := svc.NewService(oc, mc)
	if err := s.Exec(); err != nil {
		fmt.Errorf("failed to exec: %#v", err)
	}
}
