package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	userChan := make(chan GetUserResult, 1)
	groupChan := make(chan GetGroupResult, 1)
	todosChan := make(chan ListTodoResult, 1)

	go getUser(userChan)
	go getGroup(groupChan)
	go listTodo(todosChan)

	userResult := <-userChan
	groupResult := <-groupChan
	todosResult := <-todosChan
	close(userChan)
	close(groupChan)
	close(todosChan)

	var errorsList []error
	if userResult.Err != nil {
		errorsList = append(errorsList, userResult.Err)
	}
	if groupResult.Err != nil {
		errorsList = append(errorsList, groupResult.Err)
	}
	if todosResult.Err != nil {
		errorsList = append(errorsList, todosResult.Err)
	}
	if len(errorsList) > 0 {
		for _, err := range errorsList {
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	} else {
		result := Result{
			UserName:  userResult.User.Name,
			GroupName: groupResult.Group.Name,
			Contents:  todosResult.Contents(),
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
type GetUserResult struct {
	User User
	Err  error
}

func getUser(userChan chan GetUserResult) {
	time.Sleep(1 * time.Second)
	user := User{ID: 1, Name: "John"}
	// 正常系
	userChan <- GetUserResult{user, nil}
	// 異常系
	//userChan <- GetUserResult{User{}, errors.New("user not found")}
}

type Group struct {
	ID   int
	Name string
}
type GetGroupResult struct {
	Group Group
	Err   error
}

func getGroup(groupChan chan GetGroupResult) {
	time.Sleep(2 * time.Second)
	group := Group{ID: 1, Name: "GroupA"}
	// 正常系
	groupChan <- GetGroupResult{group, nil}
	// 異常系
	//groupChan <- GetGroupResult{Group{}, errors.New("group does not exist")}
}

type Todos []Todo

type Todo struct {
	ID      int
	Content string
}
type ListTodoResult struct {
	Todos Todos
	Err   error
}

func (t ListTodoResult) Contents() []string {
	contents := make([]string, len(t.Todos))
	for _, todo := range t.Todos {
		contents = append(contents, todo.Content)
	}
	return contents
}

func listTodo(todosChan chan ListTodoResult) {
	time.Sleep(4 * time.Second)
	todos := Todos{
		{ID: 1, Content: "hello"},
		{ID: 2, Content: "bye"},
	}
	// 正常系
	todosChan <- ListTodoResult{
		Todos: todos,
		Err:   nil,
	}
	// 異常系
	//todosChan <- ListTodoResult{Todos: Todos{}, Err: errors.New("todos does not exist")}
}
