package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/usecase"
	"net/http"
)

type HTTPPresenter struct {
	c echo.Context
}

func NewHTTPPresenter(c echo.Context) *HTTPPresenter {
	return &HTTPPresenter{c: c}
}

func (h *HTTPPresenter) Present(res usecase.GetUserResponse) error {
	return h.c.JSON(http.StatusOK, res.User)
}
