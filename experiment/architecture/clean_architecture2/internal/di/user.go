package di

import (
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/delivery/handler"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/infrastructure/repository"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/usecase"

	"go.uber.org/dig"
)

func User(c *dig.Container) error {
	dependencies := []any{
		repository.NewUser,
		usecase.NewGetUserInteractor,
		handler.NewUser,
	}
	for _, dependency := range dependencies {
		if err := c.Provide(dependency); err != nil {
			return err
		}
	}
	return nil
}
