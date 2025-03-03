package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/delivery/handler"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/delivery/routing"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/di"
	"log"
)

func main() {
	container := di.BuildContainer()
	if err := container.Invoke(func(e *echo.Echo, u *handler.User, dbCloseFn di.DBCloseFunc) {
		defer dbCloseFn()
		routing.User(e, u)
		if err := e.Start(":8080"); err != nil {
			log.Fatal("サーバー起動失敗:", err)
		}
	}); err != nil {
		log.Fatal(err)
	}
}
