package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sky0621/tips-go/experiment/sql_performance/mysql/v9/ec/dsn"
)

func main() {
	db, _ := sql.Open("mysql", dsn.Dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if m == nil {
		fmt.Println("no migration")
		return
	}
	if err := m.Steps(-1); err != nil {
		fmt.Println(err)
		return
	}
}
