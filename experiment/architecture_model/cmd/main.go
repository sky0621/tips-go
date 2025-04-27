package main

import (
	"github.com/sky0621/tips-go/experiment/error_handling/server"
	"log"
)

func main() {
	if err := server.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
