package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/usecase"
	"net/http"
	"strconv"
)

type User struct {
	getUserUsecase usecase.GetUserInputPort
}

func NewUser(getUserUsecase usecase.GetUserInputPort) *User {
	return &User{getUserUsecase: getUserUsecase}
}

func (u *User) GetUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	presenter := NewHTTPPresenter(c)
	req := usecase.GetUserRequest{UserID: id}
	ctx := c.Request().Context()
	if err := u.getUserUsecase.GetUser(ctx, req, presenter); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return nil
}
