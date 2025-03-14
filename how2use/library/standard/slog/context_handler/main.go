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
	slog.InfoContext(ctx, "hello world", "key1", "value1")
}
