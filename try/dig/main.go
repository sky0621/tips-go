package main

import (
	"dig/handler"
	"dig/infra"
	"dig/usecase"
	"log"

	"go.uber.org/dig"
)

func main() {
	container, err := BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	var userHandler handler.UserHandler
	if err = container.Invoke(func(u handler.UserHandler) {
		userHandler = u
	}); err != nil {
		log.Fatal(err)
	}
	_, err = userHandler.GetUser(333)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildContainer() (*dig.Container, error) {
	dependencies := []any{
		infra.NewUserRepository,
		usecase.NewUserUseCase,
	}
	container := dig.New()
	for _, dependency := range dependencies {
		if err := container.Provide(dependency); err != nil {
			return nil, err
		}
	}
	return container, nil
}
