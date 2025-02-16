package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/internal/usecase"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{todoUsecase: todoUsecase}
}

func (h *TodoHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/todos", h.CreateTodo)
	e.GET("/todos", h.ListTodos)
	e.GET("/todos/:id", h.GetTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)
}

type CreateTodoRequest struct {
	Title string `json:"title" validate:"required"`
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	req := new(CreateTodoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	todo, err := h.todoUsecase.CreateTodo(c.Request().Context(), req.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) ListTodos(c echo.Context) error {
	todos, err := h.todoUsecase.ListTodos(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
	}
	todo, err := h.todoUsecase.GetTodo(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Todoが見つかりません"})
	}
	return c.JSON(http.StatusOK, todo)
}

type UpdateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
	}
	req := new(UpdateTodoRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	todo, err := h.todoUsecase.UpdateTodo(c.Request().Context(), id, req.Title, req.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
	}
	if err := h.todoUsecase.DeleteTodo(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
