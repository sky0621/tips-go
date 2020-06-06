package svc

type Service interface {
	Exec() error
}

type service struct {
	orderAPIClient client.OrderAPIClient
	mailClient     client.MailClient
}

func NewService(orderAPIClient client.OrderAPIClient, mailClient client.MailClient) Service {
	return &service{
		orderAPIClient: orderAPIClient,
		mailClient:     mailClient,
	}
}

func (s *service) Exec() error {
	// FIXME:
	return nil
}
