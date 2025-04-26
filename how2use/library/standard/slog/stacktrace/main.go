package main

import (
	"context"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/service"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/handler"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/infra"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/log"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/usecase"
	"log/slog"
)

func main() {
	log.SetDefaultLogger()

	h := handler.NewUser(usecase.NewUser(service.NewUser(infra.NewUser())))
	_, err := h.ListUser(context.Background())
	if err != nil {
		slog.Error("failed to list user",
			slog.Any("error", err),
		)
	}
}
