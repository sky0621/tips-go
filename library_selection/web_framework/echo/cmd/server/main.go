package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) GetAppName() string {
	return "MyApp"
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		return cc.String(200, cc.GetAppName())
	})
	e.GET("/path/:id", pathParameterSample)
	e.GET("/query", queryParameterSample)
	e.POST("/users", postFormUrlEncodedSample)
	e.POST("/users2", handlingRequestSample)
	e.GET("/cookie", readCookie)
	e.POST("/cookie", writeCookie)
	e.POST("/custom_data_bind", customDataBindSample)
	e.Logger.Fatal(e.Start(":1323"))
}
