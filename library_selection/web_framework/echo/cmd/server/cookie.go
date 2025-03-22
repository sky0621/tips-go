package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const cookieKey = "appval"

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie(cookieKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, cookie)
}

type AppCookie struct {
	Value string `json:"value"`
}

func writeCookie(c echo.Context) error {
	ac := new(AppCookie)
	if err := c.Bind(ac); err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = cookieKey
	cookie.Path = "/"
	cookie.Value = ac.Value
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, "write a cookie")
}
