package domain

type User struct {
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
}
