package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.GET("/users/:id", getUserHandler)
	e.GET("/show", getShowHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
