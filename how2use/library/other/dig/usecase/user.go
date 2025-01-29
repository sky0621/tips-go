package usecase

import (
	"dig/domain"
	"fmt"
)

type User interface {
	GetUser(id int) (*domain.User, error)
	AddUser(name string) error
}

type user struct {
	r domain.UserRepository
}

func (u *user) AddUser(name string) error {
	res, err := u.r.AddUser(name)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (u *user) GetUser(id int) (*domain.User, error) {
	return u.r.GetUser(id)
}

func NewUser(r domain.UserRepository) User {
	return &user{r}
}
