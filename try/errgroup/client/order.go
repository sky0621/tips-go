package client

import (
	"context"
	"fmt"
	"time"
)

type OrderAPIClient interface {
	Request(context.Context, Order) error
}

type orderAPIClient struct {
}

func NewOrderAPIClient() OrderAPIClient {
	return &orderAPIClient{}
}

// MEMO: 「注文」情報を外部APIに渡す機能（平均 10 秒くらい時間がかかる想定）のダミー実装
func (c *orderAPIClient) Request(ctx context.Context, o Order) error {
	switch o.ID {
	case "id001":
		time.Sleep(9 * time.Second)
		fmt.Println("[Request] id001")
	case "id002":
		time.Sleep(12 * time.Second)
		fmt.Println("[Request]       id002")
	case "id003":
		//return xerrors.New("failed to Request")
		time.Sleep(10 * time.Second)
		fmt.Println("[Request]             id003")
	case "id004":
		time.Sleep(8 * time.Second)
		fmt.Println("[Request]                   id004")
	case "id005":
		time.Sleep(11 * time.Second)
		fmt.Println("[Request]                         id005")
	}
	return nil
}
