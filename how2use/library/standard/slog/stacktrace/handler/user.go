package handler

import (
	"context"
	"fmt"
	"github.com/sky0621/tips-go/how2use/library/standard/slog/stacktrace/usecase"
)

type User struct {
	u usecase.User
}

func NewUser(u usecase.User) User {
	return User{u}
}

type UserDto struct {
	ID   int
	Name string
}

func (u *User) ListUser(ctx context.Context) ([]UserDto, error) {
	res, err := u.u.ListUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("[handler] failed to list users: %w", err)
	}
	var userDtos []UserDto
	for _, user := range res {
		userDtos = append(userDtos, UserDto{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return userDtos, nil
}
