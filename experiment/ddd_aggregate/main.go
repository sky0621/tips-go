package main

import (
	"fmt"
	"log"

	"github.com/sky0621/tips-go/experiment/ddd_aggregate/order"
)

func main() {
	orderID, err := order.NewOrderID("ORD-1001")
	if err != nil {
		log.Fatal(err)
	}
	customerID, err := order.NewCustomerID("CUS-9001")
	if err != nil {
		log.Fatal(err)
	}

	o, err := order.New(orderID, customerID)
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

	if err := o.AddItem(coffeeID, 2, 450); err != nil {
		log.Fatal(err)
	}
	if err := o.AddItem(teaID, 1, 300); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("order %s total: %d\n", o.ID(), o.TotalPrice())
	for _, item := range o.Items() {
		fmt.Printf("- %s x%d = %d\n", item.ProductID(), item.Quantity(), item.Subtotal())
	}

	if err := o.Submit(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status after submit: %s\n", o.Status())

	if err := o.AddItem(coffeeID, 1, 450); err != nil {
		fmt.Printf("add after submit: %v\n", err)
	}
}
