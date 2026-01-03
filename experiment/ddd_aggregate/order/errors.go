package order

import "errors"

var (
	ErrEmptyOrderID      = errors.New("order id is required")
	ErrEmptyCustomerID   = errors.New("customer id is required")
	ErrEmptyProductID    = errors.New("product id is required")
	ErrInvalidQuantity   = errors.New("quantity must be greater than zero")
	ErrInvalidUnitPrice  = errors.New("unit price must be zero or positive")
	ErrUnitPriceMismatch = errors.New("unit price mismatch for same product")
	ErrOrderSubmitted    = errors.New("order is already submitted")
	ErrItemNotFound      = errors.New("order item not found")
	ErrNoItemsToSubmit   = errors.New("order must have at least one item to submit")
)
