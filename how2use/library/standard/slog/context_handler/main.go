package main

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"os"
)

func main() {
	slog.SetDefault(slog.New(&AppLogHandler{slog.NewJSONHandler(os.Stdout, nil)}))

	handleRequest()
}

func handleRequest() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, LogKeyTraceID, uuid.New().String())
	ctx = context.WithValue(ctx, LogKeyUserID, 101)

	getUser(ctx)
}

func getUser(ctx context.Context) {
	/*
		{
			"time":"2025-04-02T00:10:15.731512+09:00",
			"level":"INFO",
			"msg":"hello world",
			"key1":"value1",
			"trace_id":"90fcf25f-0f86-4066-b2f7-371306f66380",
			"user_id":101
		}
	*/
	slog.InfoContext(ctx, "hello world", "key1", "value1")
}
