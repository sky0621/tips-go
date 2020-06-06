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
	dbClient client.DBClient
}

func NewService(dbClient client.DBClient) Service {
	return &service{dbClient: dbClient}
}

func (s *service) Exec() error {
	orders, err := s.dbClient.CollectOrders()
	if err != nil {
		return xerrors.Errorf("failed to collect orders: %w", err)
	}

	// ★ 同時実行 goroutine 数の制御のためにチャネル用意
	semaphore := make(chan struct{}, 3)

	// ★ エラーグループ生成
	errGrp, eCtx := errgroup.WithContext(context.Background())

	fmt.Printf("START: %s\n", time.Now().Format(time.RFC3339))

	for idx, order := range orders {
		sr := subroutine{
			dbClient:       s.dbClient,
			orderAPIClient: client.NewOrderAPIClient(),
			mailClient:     client.NewMailClient(),
		}
		// [参照] https://golang.org/doc/faq#closures_and_goroutines
		o := order

		// 3個までは詰められる（でも、それ以上はチャネルが空くまで詰められずにここで待機状態となる）
		semaphore <- struct{}{}

		fmt.Printf("loop-index: %d [ID:%s][Name:%s]\n", idx, o.ID, o.Name)

		errGrp.Go(func() error {
			select {
			case <-eCtx.Done():
				return xerrors.Errorf("canceled: %#v", o)
			default:
				return sr.exec(eCtx, o, semaphore)
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
	dbClient       client.DBClient
	mailClient     client.MailClient
}

func (sr subroutine) exec(eCtx context.Context, order client.Order, semaphore chan struct{}) error {
	defer func() {
		<-semaphore // 処理後にチャネルから値を抜き出さないと、次の goroutine が起動できない
	}()

	// 「注文」情報リクエスト中であることを記録
	if err := sr.dbClient.WriteRequesting(eCtx, order); err != nil {
		return xerrors.Errorf("failed to write order requesting: %w", err)
	}
	// 外部APIを使って「注文」リクエストを飛ばす
	if err := sr.orderAPIClient.Request(eCtx, order); err != nil {
		return xerrors.Errorf("failed to request order: %w", err)
	}
	// 「注文」結果をDBに保存
	if err := sr.dbClient.SaveOrderRequestStatus(eCtx, order); err != nil {
		return xerrors.Errorf("failed to save order status: %w", err)
	}
	// 「注文」結果をメール送信
	if err := sr.mailClient.Send(eCtx, order); err != nil {
		return xerrors.Errorf("failed to send mail: %w", err)
	}
	return nil
}
