package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	contentsController "github.com/sky0621/tips-go/experiment/architecture_model/internal/content/controller"
	coursesController "github.com/sky0621/tips-go/experiment/architecture_model/internal/course/controller"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/handler/interfaces"
	"log"
)

func New() App {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	swagger, err := interfaces.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}
	oapiRequestValidator := oapimiddleware.OapiRequestValidator(swagger)
	router := e.Group("/api/v1", oapiRequestValidator)

	interfaces.RegisterHandlers(router, handler.New(contentsController.Content{}, coursesController.Course{}, nil))
	return &app{echo: e}
}

type App interface {
	Run() error
}

type app struct {
	echo *echo.Echo
}

func (a app) Run() error {
	if err := a.echo.Start(":8080"); err != nil {
		return err
	}
	return nil
}
