package main

import (
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	errorLogger := slog.NewLogLogger(handler, slog.LevelError)
	errorLogger.Println("This is an info message")
}
