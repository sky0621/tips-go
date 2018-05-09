package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteClient interface {
	Create() error
	Save()
}

type World struct {
	ID      int
	Country string
	Capital string
}

type SqliteClientImpl struct {
}

func NewSqliteClient(dbname string) SqliteClient {
	return &SqliteClientImpl{}
}

func (c *SqliteClientImpl) Create() error {
	dbfile := "./test.db"
	os.Remove(dbfile)
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))`)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`INSERT INTO "world"("country", "capital") VALUES(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec("日本", "東京"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("アメリカ", "ワシントンD.C."); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("ロシア", "モスクワ"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("イギリス", "ロンドン"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("オーストラリア", "シドニー"); err != nil {
		log.Fatal(err)
	}

	rows, qerr := db.Query(`SELECT * FROM world`)
	if qerr != nil {
		log.Fatal(qerr)
	}
	var id int
	var country string
	var capital string
	for rows.Next() {
		err = rows.Scan(&id, &country, &capital)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, " - ", country, " - ", capital)
	}

}

func (c *SqliteClientImpl) Create() error {
	dbfile := "./test.db"
	os.Remove(dbfile)
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))`)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`INSERT INTO "world"("country", "capital") VALUES(?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec("日本", "東京"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("アメリカ", "ワシントンD.C."); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("ロシア", "モスクワ"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("イギリス", "ロンドン"); err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec("オーストラリア", "シドニー"); err != nil {
		log.Fatal(err)
	}

	rows, qerr := db.Query(`SELECT * FROM world`)
	if qerr != nil {
		log.Fatal(qerr)
	}
	var id int
	var country string
	var capital string
	for rows.Next() {
		err = rows.Scan(&id, &country, &capital)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, " - ", country, " - ", capital)
	}

}
