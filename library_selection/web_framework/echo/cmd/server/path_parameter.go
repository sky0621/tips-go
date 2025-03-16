package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func pathParameterSample(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
