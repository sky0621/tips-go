package client

type MailClient interface {
	Send() error
}

type mailClient struct {
}

func NewMailClient() MailClient {
	return &mailClient{}
}

func (c *mailClient) Send() error {
	return nil
}
