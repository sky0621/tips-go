package domain

import (
	"context"
)

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *Todo) error
	GetTodos(ctx context.Context) ([]*Todo, error)
	GetTodoByID(ctx context.Context, id int64) (*Todo, error)
	UpdateTodo(ctx context.Context, todo *Todo) error
	DeleteTodo(ctx context.Context, id int64) error
}
