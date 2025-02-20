package main

import (
	"fmt"
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

func getUserHandler(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getShowHandler(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, fmt.Sprintf("team:%s - member:%s", team, member))
}
