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
		/*
			{
			  "time": "2025-05-05T12:22:20.523052+09:00",
			  "level": "ERROR",
			  "source": {
			    "function": "main.main",
			    "file": "/Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/main.go",
			    "line": 43
			  },
			  "msg": "failed to list user",
			  "error": {
			    "msg": "unexpected error",
			    "trace": [
			      "main.main /Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/main.go:17",
			      "github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/handler.(*User).ListUser /Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/handler/user.go:22",
			      "github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/usecase.User.ListUser /Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/usecase/user.go:18",
			      "github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/service.User.ListUser /Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/service/user.go:18",
			      "github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/infra.user.ListUser /Users/sky0621/work/github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/infra/user.go:19"
			    ]
			  }
			}
		*/
		slog.Error("failed to list user",
			slog.Any("error", err),
		)
	}
}
