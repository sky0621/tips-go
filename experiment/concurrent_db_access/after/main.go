package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	errChan := make(chan error, 3)

	var user User
	var group Group
	var todos Todos
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		user, err = getUser()
		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		group, err = getGroup()
		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		todos, err = listTodo()
		if err != nil {
			errChan <- err
		}
	}()

	wg.Wait()
	close(errChan)

	var errList []error
	for err := range errChan {
		errList = append(errList, err)
	}

	if len(errList) > 0 {
		for _, err := range errList {
			fmt.Println(err.Error())
		}
	} else {
		result := Result{
			UserName:  user.Name,
			GroupName: group.Name,
			Contents:  todos.Contents(),
		}
		fmt.Println(result)
	}

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
func getUserErr() (User, error) {
	time.Sleep(1 * time.Second)
	return User{}, errors.New("no user")
}

type Group struct {
	ID   int
	Name string
}

func getGroup() (Group, error) {
	time.Sleep(2 * time.Second)
	return Group{ID: 1, Name: "John"}, nil
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
func listTodoErr() (Todos, error) {
	time.Sleep(4 * time.Second)
	return nil, errors.New("no todos")
}
