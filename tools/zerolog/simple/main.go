package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//Simple Logging Example
	//For simple logging, import the global logger package github.com/rs/zerolog/log
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// f.e. {"level":"debug","time":1596760873,"message":"hello world"}
	log.Print("hello world")
}
