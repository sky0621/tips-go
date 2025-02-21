package main

import (
	"fmt"
)

func main() {
	users := Users{{ID: 1, Name: "Taro"}, {ID: 2, Name: "Jiro"}, {ID: 3, Name: "Saburo"}}
	fmt.Println(users.IDs())
	fmt.Println(users.Names())
}

type UserID int
type UserName string

type User struct {
	ID   UserID
	Name UserName
}

type Users []User

func (u Users) IDs() []UserID {
	ids := make([]UserID, len(u))
	for i, uu := range u {
		ids[i] = uu.ID
	}
	return ids
}

func (u Users) Names() []UserName {
	names := make([]UserName, len(u))
	for i, uu := range u {
		names[i] = uu.Name
	}
	return names
}
