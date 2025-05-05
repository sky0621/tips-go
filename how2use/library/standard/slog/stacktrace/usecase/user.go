package usecase

import (
	"context"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/model"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/service"
)

type User struct {
	s service.User
}

func NewUser(s service.User) User {
	return User{s}
}

func (u User) ListUser(ctx context.Context) ([]model.User, error) {
	users, err := u.s.ListUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
