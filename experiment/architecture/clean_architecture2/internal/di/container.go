package di

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/config"
	"github.com/sky0621/tips-go/experiment/architecture/clean_architecture2/internal/infrastructure/db"
	"go.uber.org/dig"
	"log"
	"time"
)

type DBCloseFunc func()

func BuildContainer() *dig.Container {
	container := dig.New()

	if err := container.Provide(func() *echo.Echo {
		return echo.New()
	}); err != nil {
		log.Fatalf("Echoインスタンス生成に失敗しました: %v", err)
	}

	if err := container.Provide(func() (*sql.DB, DBCloseFunc) {
		dbConn, err := sql.Open("mysql", config.Dsn)
		if err != nil {
			log.Fatalf("DB接続に失敗しました: %v", err)
		}

		dbConn.SetConnMaxLifetime(3 * time.Minute)
		dbConn.SetMaxOpenConns(10)
		dbConn.SetMaxIdleConns(10)
		if err = dbConn.Ping(); err != nil {
			log.Fatalf("DBのPingに失敗しました: %v", err)
		}
		return dbConn, func() {
			if err := dbConn.Close(); err != nil {
				log.Fatalf("DBのCloseに失敗しました: %v", err)
			}
		}
	}); err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}

	if err := container.Provide(func(dbConn *sql.DB) *db.Queries {
		return db.New(dbConn)
	}); err != nil {
		log.Fatalf("sqlcインスタンス生成に失敗しました: %v", err)
	}

	if err := User(container); err != nil {
		log.Fatalf("User関連インスタンス生成に失敗しました: %v", err)
	}

	return container
}
