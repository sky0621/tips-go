package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/experiment/architecture/onion_architecture/config"
	"log"
)

func main() {
	db, err := sql.Open("mysql", config.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}
}
