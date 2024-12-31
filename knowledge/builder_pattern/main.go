package main

import (
	"log"
	"os"
	"time"

	"github.com/sky0621/tips-go/knowledge/builder_pattern/server"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags|log.Lmicroseconds)

	svr := server.NewParamBuilder("localhost", 8080).
		Timeout(time.Minute).
		Logger(logger).
		Build()
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
