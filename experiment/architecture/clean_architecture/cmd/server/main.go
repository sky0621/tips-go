package main

import (
	"database/sql"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/config"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/internal/handler"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/internal/repository/mysql"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture/internal/usecase"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbConn, err := sql.Open("mysql", config.Dsn)
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}
	defer dbConn.Close()

	dbConn.SetConnMaxLifetime(3 * time.Minute)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)
	if err = dbConn.Ping(); err != nil {
		log.Fatalf("DBのPingに失敗しました: %v", err)
	}

	repo := mysql.NewTodoRepository(dbConn)
	todoUsecase := usecase.NewTodoUsecase(repo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	todoHandler.RegisterRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
