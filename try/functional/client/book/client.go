package book

// Client ...
type Client interface {
	Search(title string) []*Book
}

type client struct {
}

// NewClient ...
func NewClient() Client {
	return &client{}
}

// Search ...
func (c *client) Search(title string) []*Book {
	// FIXME
	return []*Book{}
}

// Book ...
type Book struct {
	ID       int
	Title    string
	Category int
}
