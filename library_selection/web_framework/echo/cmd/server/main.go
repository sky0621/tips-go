package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	e.GET("/path/:id", pathParameterSample)
	e.GET("/query", queryParameterSample)
	e.POST("/users", postFormUrlEncodedSample)
	e.POST("/users2", handlingRequestSample)
	e.Logger.Fatal(e.Start(":1323"))
}
