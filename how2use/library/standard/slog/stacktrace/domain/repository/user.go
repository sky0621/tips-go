package repository

import (
	"context"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/model"
)

type User interface {
	ListUser(ctx context.Context) ([]model.User, error)
}
