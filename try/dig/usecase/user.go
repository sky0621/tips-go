package usecase

import "dig/domain"

type UserUsecase interface {
	GetUser(id int) (*domain.User, error)
}

type userUsecase struct {
	r domain.UserRepository
}

func (u *userUsecase) GetUser(id int) (*domain.User, error) {
	return u.r.GetUser(id)
}

func NewUserUsecase(r domain.UserRepository) UserUsecase {
	return &userUsecase{r}
}
