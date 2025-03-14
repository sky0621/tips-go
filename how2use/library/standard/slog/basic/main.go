package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	slog.Info("slog info")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("hello world", "key1", "value1", "key2", "value2")
	logger.Debug("debug message", "dk1", "dv1", "dk2", "dv2") // この行は出力されない
	logger2 := logger.With("service", "my-service")

	slog.SetDefault(logger2)
	slog.Info("slog info2")

	logger3 := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				a.Key = "logged_at"
				a.Value = slog.StringValue(a.Value.Time().Format(time.UnixDate))
			}
			return a
		},
	}))
	logger3.Info("logger3 info")

	logger3.InfoContext(context.Background(), "logger3 info2")
}
