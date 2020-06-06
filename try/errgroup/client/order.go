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

func (c *orderAPIClient) Request(ctx context.Context, o Order) error {
	switch o.ID {
	case "id001":
		time.Sleep(350 * time.Millisecond)
		fmt.Println("[Request     ] id001")
	case "id002":
		time.Sleep(600 * time.Millisecond)
		fmt.Println("[Request     ]       id002")
	case "id003":
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("[Request     ]             id003")
	case "id004":
		time.Sleep(1800 * time.Millisecond)
		fmt.Println("[Request     ]                   id004")
	case "id005":
		time.Sleep(3000 * time.Millisecond)
		fmt.Println("[Request     ]                         id005")
	}
	return nil
}
