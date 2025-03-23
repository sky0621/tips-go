package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/library_selection/web_framework/echo/cmd/server/logger"
	"log/slog"
	"net/http"
)

func requestLoggerSample(c echo.Context) error {
	userID := c.Param("id")
	ctx := context.WithValue(c.Request().Context(), logger.LogKeyUserID, userID)
	ctx = context.WithValue(ctx, logger.LogKeyTraceID, c.Response().Header().Get(echo.HeaderXRequestID))
	slog.InfoContext(ctx, "requestLoggerSample")
	return c.String(http.StatusOK, "request logged")
}
