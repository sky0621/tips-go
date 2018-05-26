package recommend

// Client ...
type Client interface {
}

type client struct {
}

// NewClient ...
func NewClient() Client {
	return &client{}
}
