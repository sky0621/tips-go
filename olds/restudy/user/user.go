package user

type User interface {
}

func NewUser(id int, name string) User {
	return &user{
		id:   ID(id),
		name: Name{name},
	}
}

type user struct {
	id   ID
	name Name
}

type ID int

type Name struct {
	val string
}

type Age interface {
}
