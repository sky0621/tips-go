package main

import (
	"log"
	"os"
	"time"

	"github.com/sky0621/tips-go/experiment/functional_option_pattern/p01/server"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags|log.Lmicroseconds)

	svr := server.New("localhost", 8080,
		server.WithTimeout(time.Microsecond),
		server.WithLogger(logger),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
