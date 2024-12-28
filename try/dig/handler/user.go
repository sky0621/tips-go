package handler

import (
	"dig/domain"
)

type UserHandler interface {
	GetUser(id int) (*domain.User, error)
}

type userHandler struct{}

func (u *userHandler) GetUser(id int) (*domain.User, error) {
	return &domain.User{}, nil
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}
