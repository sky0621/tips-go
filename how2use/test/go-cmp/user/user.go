package user

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func GetUser(id int) (*User, error) {
	return &User{ID: id, Name: fmt.Sprintf("User%02d", id), Age: id + 5}, nil
}

func AddUser(name, id int) error {
	return nil
}
