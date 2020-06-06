package client

import (
	"context"
	"fmt"
	"time"
)

type MailClient interface {
	Send(context.Context, Order) error
}

type mailClient struct {
}

func NewMailClient() MailClient {
	return &mailClient{}
}

// MEMO: 「注文」情報メール送信機能のダミー実装
func (c *mailClient) Send(ctx context.Context, o Order) error {
	switch o.ID {
	case "id001":
		time.Sleep(4 * time.Second)
		fmt.Println("[Send   ] id001")
	case "id002":
		time.Sleep(2 * time.Second)
		fmt.Println("[Send   ]       id002")
	case "id003":
		time.Sleep(3 * time.Second)
		fmt.Println("[Send   ]             id003")
	case "id004":
		time.Sleep(2 * time.Second)
		fmt.Println("[Send   ]                   id004")
	case "id005":
		time.Sleep(5 * time.Second)
		fmt.Println("[Send   ]                         id005")
	}
	return nil
}
