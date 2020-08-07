package main

import (
	"github.com/rs/zerolog/log"
)

//{"level":"info","time":"2020-08-07T09:53:33+09:00","message":"main"}
//{"level":"info","component":"foo","time":"2020-08-07T09:53:33+09:00","message":"hello world"}
//{"level":"warn","component":"foo","time":"2020-08-07T09:54:13+09:00","message":"warn"}
//{"level":"info","time":"2020-08-07T09:53:33+09:00","message":"main2"}
func main() {
	log.Info().Msg("main")

	subLogger := log.With().
		Str("component", "foo").
		Logger()
	subLogger.Info().Msg("hello world")
	subLogger.Warn().Msg("warn")

	log.Info().Msg("main2")
}
