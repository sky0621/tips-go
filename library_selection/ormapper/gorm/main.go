package main

import (
	"flag"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/crud"

	"github.com/sky0621/tips-go/library_selection/ormapper/gorm/dsn"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	flag.Parse()
	args := flag.Args()
	command := "R"
	if len(args) > 0 {
		command = args[0]
	}

	db, err := gorm.Open(mysql.Open(dsn.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "C":
		crud.Insert(db)
	default:
		crud.Select(db)
	}
}
