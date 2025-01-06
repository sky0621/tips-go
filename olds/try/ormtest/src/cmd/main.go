package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sky0621/tips-go/try/ormtest/src/adapter"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// MEMO: ローカルでしか使わないので、ベタ書き
	dsn := "host=localhost port=22456 dbname=tips-go-try-ormtest-infra user=postgres password=yuckyjuice sslmode=disable"
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

	customerService := adapter.NewCustomerService(db)

	r := chi.NewRouter()
	r.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := customerService.Customers(r.Context())
		if err != nil {
			if _, err := w.Write([]byte(err.Error())); err != nil {
				log.Fatal(err)
			}
		}
		for _, c := range customers {
			if _, err := w.Write([]byte(fmt.Sprintf("%#+v\n", c))); err != nil {
				log.Fatal(err)
			}
		}
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
