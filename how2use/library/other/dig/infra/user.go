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
	return &domain.User{ID: id, Name: "Taro"}, nil
}
