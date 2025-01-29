package di

import (
	"dig/handler"
	"dig/infra"
	"dig/usecase"

	"go.uber.org/dig"
)

func Todo(c *dig.Container) error {
	dependencies := []any{
		infra.NewTodoRepository,
		usecase.NewTodo,
		handler.NewTodo,
	}
	for _, dependency := range dependencies {
		if err := c.Provide(dependency); err != nil {
			return err
		}
	}
	return nil
}
