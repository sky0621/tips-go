package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/dsn"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/migration"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func main() {
	sqldb, err := sql.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
	defer func(db *bun.DB) {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}(db)

	if err := migration.Migrate(db); err != nil {
		log.Fatal(err)
	}
}
