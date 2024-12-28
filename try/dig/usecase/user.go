package usecase

import "dig/domain"

type UserUseCase interface {
	GetUser(id int) (*domain.User, error)
}

type userUseCase struct {
}

func (u *userUseCase) GetUser(id int) (*domain.User, error) {
	return &domain.User{}, nil
}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}
