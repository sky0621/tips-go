package domain

import (
	"context"
	"fmt"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int) (*User, error)
	AddUser(ctx context.Context, user AddUserInput) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

type User struct {
	ID   int
	Name string
}

type AddUserInput struct {
	Name string
}

func (u *userService) GetUserByID(ctx context.Context, id int) (*User, error) {
	fmt.Printf("[GetUserByID] id: %d\n", id)
	if id == 0 {
		return nil, ErrInvalidInput(ErrorAttribute{Key: "id", Value: id})
	}
	if id == 10 {
		return nil, ErrUnauthenticated(ErrorAttribute{Key: "id", Value: id})
	}
	if id == 20 {
		return nil, ErrUnauthorized(ErrorAttribute{Key: "id", Value: id})
	}
	if id == 30 {
		return nil, ErrResourceNotFound(ErrorAttribute{Key: "id", Value: id})
	}
	if id == 40 {
		return nil, ErrUnexpectedError(ErrorAttribute{Key: "id", Value: id})
	}
	return &User{ID: id, Name: "ユーザー１"}, nil
}

func (u *userService) AddUser(ctx context.Context, user AddUserInput) error {
	fmt.Printf("[AddUser] userService: %+v\n", user)
	if user.Name == "" {
		return ErrInvalidInput(ErrorAttribute{Key: "name", Value: ""})
	}
	return nil
}
