package main

import (
	"embed"
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed static
var contents embed.FS

//go:embed static/logo.avif
var logo []byte

func main() {
	e := echo.New()
	e.GET("/", echo.WrapHandler(http.FileServerFS(contents)))
	e.GET("/logo", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/avif; charset=utf-8", logo)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
