package usecase

import "dig/domain"

type User interface {
	GetUser(id int) (*domain.User, error)
}

type user struct {
	r domain.UserRepository
}

func (u *user) GetUser(id int) (*domain.User, error) {
	return u.r.GetUser(id)
}

func NewUser(r domain.UserRepository) User {
	return &user{r}
}
