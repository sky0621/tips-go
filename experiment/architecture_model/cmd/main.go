package main

import (
	"github.com/sky0621/tips-go/experiment/architecture_model/app"
	"log"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatal(err)
	}
}
