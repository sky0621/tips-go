package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		log.Info("call /")
		return c.JSON(http.StatusOK, "Hello World!")
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	// SIGINT通知待ち
	<-ctx.Done()

	// 10秒待機（必要？）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 新規リクエスト受付を止め、既存リクエストが終わったらシャットダウン
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
