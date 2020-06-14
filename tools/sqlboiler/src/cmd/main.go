package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	/*
	 *
	 */
	dataSourceName := "dbname=boilerdb user=postgres password=localpass sslmode=false port=21340"
	db, err := sql.Open("postgres", dataSourceName)
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
