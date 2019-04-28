package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/item", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
