package di

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()
	if err := container.Provide(echo.New); err != nil {
		return nil, err
	}
	if err := User(container); err != nil {
		return nil, err
	}
	if err := Todo(container); err != nil {
		return nil, err
	}
	return container, nil
}
