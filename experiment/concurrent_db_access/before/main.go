package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	user, err := getUser()
	if err != nil {
		panic(err)
	}
	group, err := getGroup()
	if err != nil {
		panic(err)
	}
	todos, err := listTodo()
	if err != nil {
		panic(err)
	}

	result := Result{
		UserName:  user.Name,
		GroupName: group.Name,
		Contents:  todos.Contents(),
	}
	fmt.Println(result)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

type Result struct {
	UserName  string
	GroupName string
	Contents  []string
}

type User struct {
	ID   int
	Name string
}

func getUser() (User, error) {
	time.Sleep(1 * time.Second)
	return User{ID: 1, Name: "John"}, nil
}

type Group struct {
	ID   int
	Name string
}

func getGroup() (Group, error) {
	time.Sleep(2 * time.Second)
	return Group{ID: 1, Name: "GroupA"}, nil
}

type Todos []Todo

func (t Todos) Contents() []string {
	contents := make([]string, len(t))
	for _, todo := range t {
		contents = append(contents, todo.Content)
	}
	return contents
}

type Todo struct {
	ID      int
	Content string
}

func listTodo() (Todos, error) {
	time.Sleep(4 * time.Second)
	return Todos{
		{ID: 1, Content: "hello"},
		{ID: 2, Content: "bye"},
	}, nil
}
