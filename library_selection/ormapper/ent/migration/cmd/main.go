package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/dsn"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/migration"
)

func main() {
	client, err := ent.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ent.Client) {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}(client)

	if err := migration.Migrate(client); err != nil {
		log.Fatal(err)
	}
}
