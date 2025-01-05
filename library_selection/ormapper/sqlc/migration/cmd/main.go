package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/dsn"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/migration"
)

func main() {
	db, err := sql.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err := migration.Migrate(db); err != nil {
		log.Fatal(err)
	}
}
