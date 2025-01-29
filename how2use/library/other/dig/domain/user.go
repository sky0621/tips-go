package domain

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
	AddUser(name string) (*User, error)
}
