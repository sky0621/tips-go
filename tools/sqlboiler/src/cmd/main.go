package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	/*
	 * DB接続準備
	 */
	var db *sql.DB
	{
		dataSourceName := "dbname=boilerdb user=postgres password=localpass sslmode=false port=21340"
		var err error
		db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
			panic(err)
		}
		defer func() {
			if db != nil {
				if err := db.Close(); err != nil {
					panic(err)
				}
			}
		}()

		boil.DebugMode = true

		var loc *time.Location
		loc, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic(err)
		}
		boil.SetLocation(loc)
	}

}
