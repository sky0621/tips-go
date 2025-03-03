package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/delivery/handler"
)

func User(e *echo.Echo, u *handler.User) {
	e.GET("/users/:id", u.GetUser)
}
