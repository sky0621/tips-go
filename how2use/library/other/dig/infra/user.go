package infra

import (
	"dig/domain"
	"fmt"
)

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (u *userRepository) GetUser(id int) (*domain.User, error) {
	fmt.Println(id)
	switch id {
	case 1:
		return &domain.User{ID: id, Name: "Taro"}, nil
	case 2:
		return &domain.User{ID: id, Name: "Jiro"}, nil
	default:
		return nil, fmt.Errorf("user not found")
	}
}

func (u *userRepository) AddUser(name string) (*domain.User, error) {
	fmt.Println(name)
	return &domain.User{ID: 1, Name: name}, nil
}
