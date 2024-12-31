package user

import "errors"

type User struct {
	ID      int
	Name    string
	Age     int
	Hobbies []string
}

func AddUser(u User) error {
	if u.ID == 0 {
		return errors.New("user ID is zero")
	}
	// 何かしらの登録処理
	return nil
}
