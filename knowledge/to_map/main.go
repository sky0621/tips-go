package main

import "fmt"

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	userMap := toMap(users, func(u User) int {
		return u.ID
	})
	fmt.Println(userMap)
}

type User struct {
	ID   int
	Name string
	Age  int
}
