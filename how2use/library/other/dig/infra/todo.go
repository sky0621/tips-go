package infra

import (
	"dig/domain"
	"fmt"
)

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{}
}

type todoRepository struct{}

func (u *todoRepository) GetTodo(id int) (*domain.Todo, error) {
	fmt.Println(id)
	return &domain.Todo{ID: id, Content: "study golang"}, nil
}

func (u *todoRepository) AddTodo(content string) (*domain.Todo, error) {
	fmt.Println(content)
	return &domain.Todo{ID: 1, Content: content}, nil
}
