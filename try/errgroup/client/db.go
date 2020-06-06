package client

import (
	"fmt"
	"time"

	"golang.org/x/xerrors"
)

type DBClient interface {
	CollectOrders() ([]*Order, error)
}

type dbClient struct {
}

func NewDBClient() DBClient {
	return &dbClient{}
}

func (c *dbClient) CollectOrders() ([]*Order, error) {
	return []*Order{
		{ID: "id001", Name: "注文1"},
		{ID: "id002", Name: "注文2"},
		{ID: "id003", Name: "注文3"},
		{ID: "id004", Name: "注文4"},
		{ID: "id005", Name: "注文5"},
	}, nil
}

type Order struct {
	ID    string
	Name  string
	Items []Item
}

type Item struct {
	ID   string
	Name string
}

func (o Order) CollectItems() error {
	switch o.ID {
	case "id001":
		o.Items = []Item{
			{ID: "id00a", Name: "商品1"},
			{ID: "id00b", Name: "商品2"},
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("[CollectItems] id001")
	case "id002":
		o.Items = []Item{
			{ID: "id00b", Name: "商品2"},
		}
		time.Sleep(1500 * time.Millisecond)
		fmt.Println("[CollectItems]       id002")
	case "id003":
		o.Items = []Item{
			{ID: "id00c", Name: "商品3"},
			{ID: "id00d", Name: "商品4"},
			{ID: "id00e", Name: "商品5"},
		}
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("[CollectItems]             id003")
	case "id004":
		o.Items = []Item{
			{ID: "id00a", Name: "商品1"},
			{ID: "id00d", Name: "商品4"},
			{ID: "id00e", Name: "商品5"},
		}
		time.Sleep(800 * time.Millisecond)
		fmt.Println("[CollectItems]                   id004")
	case "id005":
		o.Items = []Item{
			{ID: "id00c", Name: "商品3"},
		}
		fmt.Println("[CollectItems]                         id005")
		return xerrors.Errorf("failed to collect with id005")
	}
	return nil
}
