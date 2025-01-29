package handler

import (
	"dig/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Todo interface {
	GetTodo(c echo.Context) error
	AddTodo(c echo.Context) error
}

type todo struct {
	u usecase.Todo
}

func (u *todo) GetTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid todo ID"})
	}
	todo, err := u.u.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

func (u *todo) AddTodo(c echo.Context) error {
	var request struct {
		Content string `json:"content"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	if err := u.u.AddTodo(request.Content); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add todo"})
	}
	return c.String(http.StatusCreated, "Todo added successfully")
}

func NewTodo(u usecase.Todo) Todo {
	return &todo{u}
}
