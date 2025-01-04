package main

import (
	"context"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/crud"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/dsn"
	"github.com/sky0621/tips-go/library_selection/ormapper/ent/ent"
)

func main() {
	flag.Parse()
	args := flag.Args()
	command := "R"
	if len(args) > 0 {
		command = args[0]
	}

	client, err := ent.Open("mysql", dsn.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *ent.Client) {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}(client)

	ctx := context.Background()

	switch command {
	case "C":
		crud.Insert(ctx, client)
	case "U":
		//crud.Update(ctx, client)
	case "D":
		//crud.Delete(ctx, client)
	case "T":
		//crud.Transaction(ctx, client)
	default:
		crud.Select(ctx, client)
	}
}
