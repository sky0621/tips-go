package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func main() {
	// 025/04/02 00:07:12 INFO slog info
	slog.Info("slog info")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // INFOレベル以上のログじゃないと出力されなくする
	}))

	/*
		{"time":"2025-04-02T00:07:37.020324+09:00","level":"INFO","msg":"hello world","key1":"value1","key2":"value2"}
	*/
	logger.Info("hello world", "key1", "value1", "key2", "value2")

	logger.Debug("debug message", "dk1", "dv1", "dk2", "dv2") // この行は出力されない

	logger2 := logger.With("service", "my-service")
	slog.SetDefault(logger2)
	/*
		{"time":"2025-04-02T00:07:37.020328+09:00","level":"INFO","msg":"slog info2","service":"my-service"}
	*/
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
	/*
		{"logged_at":"Wed Apr  2 00:07:37 JST 2025","level":"INFO","msg":"logger3 info"}
	*/
	logger3.Info("logger3 info")

	/*
		{"logged_at":"Wed Apr  2 00:07:37 JST 2025","level":"INFO","msg":"logger3 info2"}
	*/
	logger3.InfoContext(context.Background(), "logger3 info2")
}
