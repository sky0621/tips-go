package main

import (
	"fmt"
	"strings"
)

type User struct {
	ID       string
	Name     string
	LastName *string
	Friends  []User
}

func userToString(u User) string {
	ln := ""
	if u.LastName != nil {
		ln = *u.LastName
	}
	bf := strings.Builder{}
	bf.WriteString("ID:")
	bf.WriteString(u.ID)
	bf.WriteString(", Name:")
	bf.WriteString(u.Name)
	bf.WriteString(", LastName:")
	bf.WriteString(ln)
	return bf.String()
}

func (u User) String() string {
	bf := strings.Builder{}
	bf.WriteString("{")
	bf.WriteString(userToString(u))
	if len(u.Friends) > 0 {
		bf.WriteString(", [")
	}
	for idx, friend := range u.Friends {
		if idx != 0 {
			bf.WriteString(", ")
		}
		bf.WriteString("{" + userToString(friend) + "}")
	}
	if len(u.Friends) > 0 {
		bf.WriteString("]")
	}
	bf.WriteString("}")
	return bf.String()
}

func main() {
	myLast := "佐藤"
	kLast := "高橋"
	user1 := User{
		Name:     "太郎",
		LastName: &myLast,
		ID:       "100",
		Friends: []User{
			{
				Name:     "花子",
				LastName: nil,
				ID:       "200",
			},
			{
				Name:     "和男",
				LastName: &kLast,
				ID:       "300",
			},
		},
	}

	user2 := user1
	fmt.Println(user2)

	*user1.LastName = "渡辺"
	fmt.Println(user2)
}

func copyUser(u User) User {
	return User{
		ID:       "",
		Name:     "",
		LastName: nil,
		Friends:  nil,
	}
}
