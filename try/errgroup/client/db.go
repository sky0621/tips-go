package client

import (
	"context"
	"fmt"
	"time"
)

type DBClient interface {
	CollectOrders() ([]Order, error)
	WriteRequesting(context.Context, Order) error
	SaveOrderRequestStatus(context.Context, Order) error
}

type dbClient struct {
}

func NewDBClient() DBClient {
	return &dbClient{}
}

// MEMO: 「注文」情報取得DB検索機能のダミー実装
func (c *dbClient) CollectOrders() ([]Order, error) {
	return []Order{
		{ID: "id001", Name: "注文1"},
		{ID: "id002", Name: "注文2"},
		{ID: "id003", Name: "注文3"},
		{ID: "id004", Name: "注文4"},
		{ID: "id005", Name: "注文5"},
	}, nil
}

type Order struct {
	ID   string
	Name string
}

// MEMO: 「注文」情報リクエスト中である旨を記録する機能のダミー実装
func (c *dbClient) WriteRequesting(ctx context.Context, o Order) error {
	switch o.ID {
	case "id001":
		time.Sleep(50 * time.Millisecond)
		fmt.Println("[Write  ] id001")
	case "id002":
		time.Sleep(30 * time.Millisecond)
		fmt.Println("[Write  ]       id002")
	case "id003":
		time.Sleep(80 * time.Millisecond)
		fmt.Println("[Write  ]             id003")
	case "id004":
		time.Sleep(100 * time.Millisecond)
		fmt.Println("[Write  ]                   id004")
	case "id005":
		time.Sleep(65 * time.Millisecond)
		fmt.Println("[Write  ]                         id005")
	}
	return nil
}

// MEMO: 「注文」情報のリクエスト結果保存機能のダミー実装
func (c *dbClient) SaveOrderRequestStatus(ctx context.Context, o Order) error {
	switch o.ID {
	case "id001":
		time.Sleep(25 * time.Millisecond)
		fmt.Println("[Save   ] id001")
	case "id002":
		time.Sleep(60 * time.Millisecond)
		fmt.Println("[Save   ]       id002")
	case "id003":
		time.Sleep(40 * time.Millisecond)
		fmt.Println("[Save   ]             id003")
	case "id004":
		time.Sleep(55 * time.Millisecond)
		fmt.Println("[Save   ]                   id004")
	case "id005":
		time.Sleep(35 * time.Millisecond)
		fmt.Println("[Save   ]                         id005")
	}
	return nil
}
