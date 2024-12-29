package main

import (
	"dig/handler"
	"dig/infra"
	"dig/usecase"
	"fmt"
	"log"

	"go.uber.org/dig"
)

func main() {
	container, err := BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	if err = container.Invoke(getUser); err != nil {
		log.Fatal(err)
	}
}

func getUser(u handler.UserHandler) {
	user, err := u.GetUser(333)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func BuildContainer() (*dig.Container, error) {
	dependencies := []any{
		infra.NewUserRepository,
		usecase.NewUserUsecase,
		handler.NewUserHandler,
	}
	container := dig.New()
	for _, dependency := range dependencies {
		if err := container.Provide(dependency); err != nil {
			return nil, err
		}
	}
	return container, nil
}
