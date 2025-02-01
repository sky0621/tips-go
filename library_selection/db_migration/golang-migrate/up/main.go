package main

import (
	"database/sql"
	"fmt"

	"github.com/sky0621/tips-go/library_selection/db_migration/golang-migrate/dsn"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	if err := m.Up(); err != nil {
		fmt.Println(err)
		return
	}
}
