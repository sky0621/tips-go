package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.LogAttrs(
		ctx,
		slog.LevelInfo,
		"This is an info message",
		slog.Group("group",
			slog.String("key1", "value1"),
			slog.String("key2", "value2"),
			slog.Group("sub-group",
				slog.String("sub-key1", "sub-value1"),
				slog.String("sub-key2", "sub-value2"),
				slog.Group("sub-sub-group",
					slog.String("sub-sub-key1", "sub-sub-value1"),
				),
			),
		),
	)
}
