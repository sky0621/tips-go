package usecase

import (
	"dig/domain"
	"fmt"
)

type Todo interface {
	GetTodo(id int) (*domain.Todo, error)
	AddTodo(name string) error
}

type todo struct {
	r domain.TodoRepository
}

func (u *todo) AddTodo(name string) error {
	res, err := u.r.AddTodo(name)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (u *todo) GetTodo(id int) (*domain.Todo, error) {
	return u.r.GetTodo(id)
}

func NewTodo(r domain.TodoRepository) Todo {
	return &todo{r}
}
