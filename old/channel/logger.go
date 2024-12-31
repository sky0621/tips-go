package main

import (
	"io"
	"log"
	"os"
)

func setupLog() (*os.File, error) {
	logfile, err := os.OpenFile("./channel.log_with_caller_line", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	return logfile, nil
}
