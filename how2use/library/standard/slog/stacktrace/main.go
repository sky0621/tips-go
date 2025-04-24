package main

import (
	"context"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/service"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/handler"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/infra"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/usecase"
	"log/slog"
	"os"
	"runtime/debug"
)

func main() {
	h := handler.NewUser(usecase.NewUser(service.NewUser(infra.NewUser())))

	ctx := context.Background()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))

	_, err := h.ListUser(ctx)
	if err != nil {
		logger.Error("failed to list user",
			slog.String("error", err.Error()),
			slog.String("stacktrace", string(debug.Stack())),
		)
	}
}
