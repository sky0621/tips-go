package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Info("This is an info message")

	bi, _ := debug.ReadBuildInfo()

	childLogger := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", bi.GoVersion),
		),
	)

	childLogger.Info("This is an info message2")
}
