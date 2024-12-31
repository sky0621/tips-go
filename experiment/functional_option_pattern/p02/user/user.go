package user

import (
	"errors"
	"strings"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Email   string
	Hobbies []string
}

func AddUser(u User) error {
	if u.ID == 0 {
		return errors.New("user ID is zero")
	}
	if u.Name == "" {
		return errors.New("user name is empty")
	}
	if u.Age == 0 {
		return errors.New("user age is zero")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email")
	}
	// 何かしらの登録処理
	return nil
}
