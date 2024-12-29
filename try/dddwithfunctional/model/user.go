package model

type User struct {
	id    *UserID
	name  *UserName
	age   *Age
	email *Email
}

func NewUser(id *UserID, name *UserName, age *Age, email *Email) (*User, error) {
	return &User{
		id:    id,
		name:  name,
		age:   age,
		email: email,
	}, nil
}

func (u *User) ID() *UserID {
	return u.id
}

func (u *User) UserName() *UserName {
	return u.name
}

func (u *User) Age() *Age {
	return u.age
}

func (u *User) Email() *Email {
	return u.email
}

func (u *User) ChangeName(name *UserName) (*User, error) {
	return &User{
		id:    u.id,
		name:  name,
		age:   u.age,
		email: u.email,
	}, nil
}

func (u *User) ChangeEmail(newEmail *Email) (*User, error) {
	return &User{
		id:    u.id,
		name:  u.name,
		age:   u.age,
		email: newEmail,
	}, nil
}
