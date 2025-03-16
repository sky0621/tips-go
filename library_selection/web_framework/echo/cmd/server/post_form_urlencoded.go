package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func postFormUrlEncodedSample(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	return c.String(http.StatusOK, "Hello, "+name+", age: "+age)
}
