package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/crud"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/dsn"
	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func main() {
	flag.Parse()
	args := flag.Args()
	command := "R"
	if len(args) > 0 {
		command = args[0]
	}

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

	q := infra.New(db)

	ctx := context.Background()

	switch command {
	case "C":
		crud.Insert(ctx, q)
	case "U":
		crud.Update(ctx, q)
	case "D":
		crud.Delete(ctx, q)
	case "T":
		crud.Transaction(ctx, db)
	case "V":
		crud.Various(ctx, q)
	case "B":
		crud.InsertBatch(ctx, q)
	case "CM":
		crud.Complex(ctx, q)
	case "OS":
		crud.OtherSetup(ctx, q)
	case "O":
		crud.Other(ctx, q)
	case "S":
		crud.Search(ctx, q)
	case "SCS":
		crud.SchoolSetup(ctx, q)
	case "SC":
		crud.School(ctx, q)
	default:
		crud.Select(ctx, q)
	}
}
