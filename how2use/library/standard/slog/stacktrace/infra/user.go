package infra

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/model"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/repository"
)

var _ repository.User = (*user)(nil)

type user struct{}

func NewUser() repository.User {
	return &user{}
}

func (u user) ListUser(ctx context.Context) ([]model.User, error) {
	e := errors.New("unexpected error")
	return []model.User{}, e
}
