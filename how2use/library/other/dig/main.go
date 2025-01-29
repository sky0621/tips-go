package main

import (
	"dig/di"
	"dig/handler"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	container, err := di.BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	if err = container.Invoke(func(e *echo.Echo, u handler.User, t handler.Todo) {
		user(e, u)
		todo(e, t)
		e.Logger.Fatal(e.Start(":8080"))
	}); err != nil {
		log.Fatal(err)
	}
}

func user(e *echo.Echo, u handler.User) {
	e.GET("/users/:id", u.GetUser)
	e.POST("/users", u.AddUser)
}

func todo(e *echo.Echo, t handler.Todo) {
	e.GET("/todos/:id", t.GetTodo)
	e.POST("/todos", t.AddTodo)
}
