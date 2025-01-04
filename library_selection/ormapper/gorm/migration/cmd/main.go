package main

import (
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/dsn"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(dsn.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = migration.Migrate(db); err != nil {
		log.Fatal(err)
	}
}
