package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/config"
	"github.com/sky0621/tips-go/experiment/architecture_model/internal/shared/rdb"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env ファイルが見つかりませんでした")
	}

	var up, down bool
	flag.BoolVar(&up, "up", false, "全マイグレーション適用")
	flag.BoolVar(&down, "down", false, "全マイグレーションロールバック")
	flag.Parse()

	ctx := context.Background()
	cfg := config.NewConfig()
	sqlDB, err := rdb.NewDB(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if driver == nil {
		fmt.Println("driver is nil")
		return
	}
	m, _ := migrate.NewWithDatabaseInstance(
		"file://schema/db",
		"mysql",
		driver,
	)

	if m == nil {
		fmt.Println("no migration")
		return
	}
	if up {
		if err := m.Up(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("migration up success")
		return
	}
	if down {
		if err := m.Down(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("migration down success")
		return
	}
	fmt.Println("no order")
}
