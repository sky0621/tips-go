package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func requestIDSample(c echo.Context) error {
	reqID := c.Response().Header().Get(echo.HeaderXRequestID)
	return c.String(http.StatusOK, reqID)
}
