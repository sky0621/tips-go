package service

import (
	"context"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/model"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/domain/repository"
)

type User struct {
	r repository.User
}

func NewUser(r repository.User) User {
	return User{r}
}

func (u User) ListUser(ctx context.Context) ([]model.User, error) {
	users, err := u.r.ListUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
