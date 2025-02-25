package service

import (
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/api"
	"time"
)

func NewTodo() *Todo {
	return &Todo{todos: []api.Todo{}}
}

type Todo struct {
	// TODO: 本来ならRepositoryの依存等
	todos []api.Todo
}

func (t *Todo) CreateTodo(todoInput api.TodoCreateInput) error {
	t.todos = append(t.todos, api.Todo{
		TodoID:    len(t.todos) + 1,
		Title:     todoInput.Title,
		Content:   todoInput.Content,
		CreatedAt: time.Now(),
	})
	return nil
}

func (t *Todo) ListTodo() ([]api.Todo, error) {
	return t.todos, nil
}

func (t *Todo) GetTodoById(id api.TodoID) (api.Todo, error) {
	for _, todo := range t.todos {
		if todo.TodoID == id {
			return todo, nil
		}
	}
	return api.Todo{}, nil
}
