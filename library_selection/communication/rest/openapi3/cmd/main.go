package main

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/api"
	"github.com/sky0621/tips-go/library_selection/communication/rest/openapi3/internal/server"
	"log"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	swagger.Servers = nil

	svr := server.New()
	e := echo.New()
	e.Use(echomiddleware.LoggerWithConfig(echomiddleware.LoggerConfig{}))
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{}))
	e.Use(echomiddleware.RequestID())
	e.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(e, svr)
	e.Logger.Fatal(e.Start(":8080"))
}
