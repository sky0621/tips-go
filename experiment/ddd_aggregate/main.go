package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sky0621/tips-go/experiment/ddd_aggregate/order"
	"github.com/sky0621/tips-go/experiment/ddd_aggregate/rdb"
	"github.com/sky0621/tips-go/experiment/ddd_aggregate/repository"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cfg, err := rdb.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewOrderRepository(db)

	orderID, err := order.NewOrderID("ORD-1001")
	if err != nil {
		log.Fatal(err)
	}
	customerID, err := order.NewCustomerID("CUS-9001")
	if err != nil {
		log.Fatal(err)
	}

	aggregate, err := order.New(orderID, customerID)
	if err != nil {
		log.Fatal(err)
	}

	coffeeID, err := order.NewProductID("COFFEE")
	if err != nil {
		log.Fatal(err)
	}
	teaID, err := order.NewProductID("TEA")
	if err != nil {
		log.Fatal(err)
	}

	if err := aggregate.AddItem(coffeeID, 2, 450); err != nil {
		log.Fatal(err)
	}
	if err := aggregate.AddItem(teaID, 1, 300); err != nil {
		log.Fatal(err)
	}
	if err := aggregate.Submit(); err != nil {
		log.Fatal(err)
	}

	if err := repo.Save(ctx, aggregate); err != nil {
		log.Fatal(err)
	}

	loaded, err := repo.FindByID(ctx, orderID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("order %s status: %s total: %d\n", loaded.ID(), loaded.Status(), loaded.TotalPrice())
	for _, item := range loaded.Items() {
		fmt.Printf("- %s x%d = %d\n", item.ProductID(), item.Quantity(), item.Subtotal())
	}
}
