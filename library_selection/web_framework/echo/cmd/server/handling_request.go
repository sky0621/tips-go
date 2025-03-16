package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func postUsersHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(&u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
