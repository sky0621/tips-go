package client

type OrderAPIClient interface {
	Collect() ([]*Order, error)
}

type orderAPIClient struct {
}

func NewOrderAPIClient() OrderAPIClient {
	return &orderAPIClient{}
}

func (c *orderAPIClient) Collect() ([]*Order, error) {
	// FIXME:
	return nil, nil
}

type Order struct {
	ID   string
	Name string
}
