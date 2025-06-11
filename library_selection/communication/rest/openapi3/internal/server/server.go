package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/api"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/internal/service"
	"net/http"
)

var _ api.ServerInterface = (*Server)(nil)

func New() *Server {
	return &Server{todoService: service.NewTodo()}
}

type Server struct {
	todoService *service.Todo
}

func (s Server) GetTodos(ctx echo.Context, params api.GetTodosParams) error {
	// TODO: 本来ならserviceに指定して件数を絞る操作をするが省略
	fmt.Println(params.Limit)
	requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
	fmt.Println(requestID)

	todos, err := s.todoService.ListTodo()
	if err != nil {
		if err := ctx.JSON(http.StatusInternalServerError, err); err != nil {
			return err
		}
	}
	if err := ctx.JSON(http.StatusOK, todos); err != nil {
		return err
	}
	return nil
}

func (s Server) PostTodos(ctx echo.Context) error {
	var todoInput api.TodoCreateInput
	if err := ctx.Bind(&todoInput); err != nil {
		if err := ctx.JSON(http.StatusBadRequest, err); err != nil {
			return err
		}
	}
	err := s.todoService.CreateTodo(todoInput)
	if err != nil {
		if err := ctx.JSON(http.StatusInternalServerError, err); err != nil {
			return err
		}
	}
	if err := ctx.JSON(http.StatusCreated, nil); err != nil {
		return err
	}
	return nil
}

func (s Server) GetTodosTodoId(ctx echo.Context, todoId api.TodoID) error {
	todo, err := s.todoService.GetTodoById(todoId)
	if err != nil {
		if err := ctx.JSON(http.StatusInternalServerError, err); err != nil {
			return err
		}
	}
	if err := ctx.JSON(http.StatusOK, todo); err != nil {
		return err
	}
	return nil
}
