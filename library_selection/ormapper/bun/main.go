package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/crud"
	"github.com/sky0621/tips-go/library_selection/ormapper/bun/dsn"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	flag.Parse()
	args := flag.Args()
	command := "R"
	if len(args) > 0 {
		command = args[0]
	}

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

	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		),
	)

	ctx := context.Background()

	switch command {
	case "C":
		crud.Insert(ctx, db)
	case "U":
		crud.Update(ctx, db)
	case "D":
		crud.Delete(ctx, db)
	case "T":
		crud.Transaction(ctx, db)
	case "V":
		crud.Various(ctx, db)
	default:
		crud.Select(ctx, db)
	}
}
