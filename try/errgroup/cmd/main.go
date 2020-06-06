package main

import (
	svc "github.com/sky0621/tips-go/try/errgroup"

	"github.com/sky0621/tips-go/try/errgroup/client"
)

func main() {
	dc := client.NewDBClient()
	oc := client.NewOrderAPIClient()
	mc := client.NewMailClient()
	s := svc.NewService(dc, oc, mc)
	if err := s.Exec(); err != nil {
		panic(err)
	}
}
