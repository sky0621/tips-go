package handler

import (
	"dig/domain"
	"dig/usecase"
)

type User interface {
	GetUser(id int) (*domain.User, error)
}

type user struct {
	u usecase.User
}

func (u *user) GetUser(id int) (*domain.User, error) {
	return u.u.GetUser(id)
}

func NewUser(u usecase.User) User {
	return &user{u}
}
