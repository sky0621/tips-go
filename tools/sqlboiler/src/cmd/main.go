package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sky0621/tips-go/try/sqlboiler/src/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	dbx, err := setupDB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if dbx != nil {
			if err := dbx.Close(); err != nil {
				panic(err)
			}
		}
	}()

	if err := exec(context.Background(), dbx); err != nil {
		panic(err)
	}
}

// DB接続準備
func setupDB() (*sqlx.DB, error) {
	dataSourceName := "dbname=boilerdb user=postgres password=localpass sslmode=disable port=21340"
	dbx, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	boil.SetLocation(loc)

	return dbx, nil
}

func exec(ctx context.Context, dbx *sqlx.DB) error {
	tx, err := dbx.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if tx != nil {
			if err := tx.Commit(); err != nil {
				panic(err)
			}
		}
	}()

	sql := `UPDATE reviews SET reviews_id=nextval('reviews_id_seq'), result=$1 WHERE content_id=$2 AND seller_id=$3`
	res, err := tx.Exec(sql, 3, 1, 1)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", res)

	qm.SQL("UPDATE reviews SET reviews_id=nextval('reviews_id_seq'), result=$1 WHERE content_id=$2 AND seller_id=$3",
		15, 1, 1)

	review, err := models.Reviews(
		models.ReviewWhere.ContentID.EQ(1),
		models.ReviewWhere.SellerID.EQ(1),
	).One(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", review)

	return nil
}
