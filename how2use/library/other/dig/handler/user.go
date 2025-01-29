package handler

import (
	"dig/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User interface {
	GetUser(c echo.Context) error
	AddUser(c echo.Context) error
}

type user struct {
	u usecase.User
}

func (u *user) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}
	user, err := u.u.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (u *user) AddUser(c echo.Context) error {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	if err := u.u.AddUser(request.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add user"})
	}
	return c.String(http.StatusCreated, "User added successfully")
}

func NewUser(u usecase.User) User {
	return &user{u}
}
