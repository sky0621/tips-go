package sample

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func AddUser(u User) error {
	fmt.Println(u)
	return nil
}
