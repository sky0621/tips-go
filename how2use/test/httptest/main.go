package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func rootHandler(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
