package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// MEMO: ローカルでしか使わないので、ベタ書き
	dsn := "host=localhost port=23456 dbname=tips-go-try-dbconn-infra user=postgres password=yuckyjuice sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if db != nil {
			// It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
			if err := db.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}()
	log.Println("Yes")

	db.SetMaxOpenConns(80)
	db.SetMaxIdleConns(80)
	db.SetConnMaxLifetime(10 * time.Second)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			st := db.Stats()
			log.Printf("{MaxOpenConnections: %d, OpenConnections: %d, InUse: %d, Idle: %d}",
				st.MaxOpenConnections, st.OpenConnections, st.InUse, st.Idle)
		}
	}()

	r := chi.NewRouter()
	r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Queryx("SELECT * FROM customer")
		if err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Fatal(err)
			}
		}
		if rows == nil {
			log.Fatal("rows == nil")
		}
		for rows.Next() {
			var c Customer
			if err := rows.StructScan(&c); err != nil {
				log.Fatal(err)
			}
			if _, err := w.Write([]byte(fmt.Sprintf("%#+v", c))); err != nil {
				log.Fatal(err)
			}
		}
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

type Customer struct {
	ID   int64
	Name string
	Age  int64
}
