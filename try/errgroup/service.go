package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/sky0621/tips-go/try/errgroup/client"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
)

type Service interface {
	Exec() error
}

type service struct {
	dbClient       client.DBClient
	orderAPIClient client.OrderAPIClient
	mailClient     client.MailClient
}

func NewService(dbClient client.DBClient, orderAPIClient client.OrderAPIClient, mailClient client.MailClient) Service {
	return &service{
		dbClient:       dbClient,
		orderAPIClient: orderAPIClient,
		mailClient:     mailClient,
	}
}

func (s *service) Exec() error {
	orders, err := s.dbClient.CollectOrders()
	if err != nil {
		return xerrors.Errorf("failed to collect orders: %w", err)
	}

	errGrp, eCtx := errgroup.WithContext(context.Background())

	fmt.Printf("START: %s\n", time.Now().Format(time.RFC3339))

	for _, order := range orders {
		sr := subroutine{
			orderAPIClient: s.orderAPIClient,
			mailClient:     s.mailClient,
		}
		o := *order

		errGrp.Go(func() error {
			select {
			case <-eCtx.Done():
				return xerrors.Errorf("canceled: %#v", o)
			default:
				return sr.exec(eCtx, o)
			}
		})
	}

	if err := errGrp.Wait(); err != nil {
		return xerrors.Errorf("failed to exec goroutine: %w", err)
	}

	fmt.Printf("END  : %s\n", time.Now().Format(time.RFC3339))

	return nil
}

type subroutine struct {
	orderAPIClient client.OrderAPIClient
	mailClient     client.MailClient
}

func (sr subroutine) exec(eCtx context.Context, order client.Order) error {
	if err := order.CollectItems(); err != nil {
		return xerrors.Errorf("failed to collect items: %w", err)
	}

	if err := sr.orderAPIClient.Request(eCtx, order); err != nil {
		return xerrors.Errorf("failed to request order: %w", err)
	}

	if err := sr.mailClient.Send(eCtx, order); err != nil {
		return xerrors.Errorf("failed to send mail: %w", err)
	}

	return nil
}
