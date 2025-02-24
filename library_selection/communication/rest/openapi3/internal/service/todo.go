package service

import (
	"fmt"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/api"
	"time"
)

func NewTodo() *Todo {
	return &Todo{}
}

type Todo struct {
	// TODO: 本来ならRepositoryの依存等
	todos []api.Todo
}

func (t *Todo) CreateTodo(todoInput api.TodoCreateInput) {
	fmt.Println(todoInput)
	t.todos = append(t.todos, api.Todo{
		Title:     todoInput.Title,
		Content:   todoInput.Content,
		CreatedAt: time.Now(),
	})
}

func (t *Todo) ListTodo() ([]api.Todo, error) {
	return t.todos, nil
}
