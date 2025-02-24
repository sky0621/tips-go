package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/api"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/internal/service"
)

var _ api.ServerInterface = (*Server)(nil)

func NewServer() *Server {
	return &Server{todoService: service.NewTodo()}
}

type Server struct {
	todoService *service.Todo
}

func (s Server) GetTodos(ctx echo.Context) error {
	return s.todoService.ListTodo()
}

func (s Server) PostTodos(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (s Server) GetTodosTodoId(ctx echo.Context, todoId api.TodoID) error {
	// TODO implement me
	panic("implement me")
}
