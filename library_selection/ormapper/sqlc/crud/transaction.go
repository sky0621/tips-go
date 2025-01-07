package crud

import (
	"context"
	"database/sql"
	"log"

	"github.com/sky0621/tips-go/library_selection/ormapper/sqlc/infra"
)

func Transaction(ctx context.Context, db *sql.DB) {
	q := infra.New(db)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	qtx := q.WithTx(tx)
	_, err = qtx.CreateUser(ctx, "Bob")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	//_ = tx.Rollback()

	_, err = qtx.CreateUser(ctx, "Ken")
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
