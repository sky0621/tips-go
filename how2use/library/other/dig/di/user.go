package di

import (
	"dig/handler"
	"dig/infra"
	"dig/usecase"

	"go.uber.org/dig"
)

func User(c *dig.Container) error {
	dependencies := []any{
		infra.NewUserRepository,
		usecase.NewUser,
		handler.NewUser,
	}
	for _, dependency := range dependencies {
		if err := c.Provide(dependency); err != nil {
			return err
		}
	}
	return nil
}
