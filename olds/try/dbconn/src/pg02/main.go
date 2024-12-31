package main

import (
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// MEMO: ローカルでしか使わないので、ベタ書き
	dsn := "host=localhost port=23456 dbname=tips-go-try-dbconn-db user=postgres password=yuckyjuice sslmode=disable"
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

	wg := new(sync.WaitGroup)

	for i := 0; i < 800; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rows, err := db.Queryx("SELECT * FROM customer")
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
			if rows == nil {
				log.Fatal("rows == nil")
			}
			//var cs []Customer
			for rows.Next() {
				var c Customer
				if err := rows.StructScan(&c); err != nil {
					log.Fatal(err)
				}
				//cs = append(cs, c)
				//log_with_caller_line.Printf("%#+v", c)
			}
			//log_with_caller_line.Println(len(cs))
		}(wg)
	}
	wg.Wait()
	time.Sleep(3 * time.Second)
	log.Println("Done")
}

type Customer struct {
	ID   int64
	Name string
	Age  int64
}
