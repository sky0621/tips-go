package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApp() App {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	return &app{echo: e}
}

type App interface {
	Run() error
}

type app struct {
	echo *echo.Echo
}

func (a app) Run() error {
	if err := a.echo.Start(":8080"); err != nil {
		return err
	}
	return nil
}
