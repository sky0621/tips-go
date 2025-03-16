package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	e.GET("/", rootHandler)
	e.GET("/users/:id", getUserHandler)
	e.GET("/show", getShowHandler)
	e.POST("/users", postHandler)
	e.POST("/users2", postUsersHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
