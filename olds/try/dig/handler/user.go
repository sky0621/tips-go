package handler

import (
	"dig/domain"
	"dig/usecase"
)

type UserHandler interface {
	GetUser(id int) (*domain.User, error)
}

type userHandler struct {
	u usecase.UserUsecase
}

func (u *userHandler) GetUser(id int) (*domain.User, error) {
	return u.u.GetUser(id)
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{u}
}
