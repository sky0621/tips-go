package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("[before]")
	printAlloc()
	num := 1_000_000
	users := make([]User, num)
	for i := 0; i < num; i++ {
		users[i] = newUser(i+1, fmt.Sprintf("User%d", i+1), 20)
	}
	for _, user := range users {
		if user.ID == 100 {
			fmt.Println(user.Name)
		}
	}
	fmt.Println("[after create users]")
	printAlloc()
	users = nil
	runtime.GC()
	fmt.Println("[after set users nil]")
	printAlloc()
}

type User struct {
	ID   int
	Name string
	Age  int
}

func newUser(id int, name string, age int) User {
	return User{ID: id, Name: name, Age: age}
}
