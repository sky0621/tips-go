package order

import "strings"

type OrderID string
type CustomerID string
type ProductID string

type Status string

const (
	StatusDraft     Status = "draft"
	StatusSubmitted Status = "submitted"
)

type ItemSnapshot struct {
	ProductID ProductID
	Quantity  int
	UnitPrice int
}

type Order struct {
	id         OrderID
	customerID CustomerID
	status     Status
	items      map[ProductID]*OrderItem
}

func NewOrderID(value string) (OrderID, error) {
	id := OrderID(strings.TrimSpace(value))
	if err := validateOrderID(id); err != nil {
		return "", err
	}
	return id, nil
}

func NewCustomerID(value string) (CustomerID, error) {
	id := CustomerID(strings.TrimSpace(value))
	if err := validateCustomerID(id); err != nil {
		return "", err
	}
	return id, nil
}

func NewProductID(value string) (ProductID, error) {
	id := ProductID(strings.TrimSpace(value))
	if err := validateProductID(id); err != nil {
		return "", err
	}
	return id, nil
}

func ParseStatus(value string) (Status, error) {
	status := Status(strings.TrimSpace(value))
	if err := validateStatus(status); err != nil {
		return "", err
	}
	return status, nil
}

func New(id OrderID, customerID CustomerID) (*Order, error) {
	if err := validateOrderID(id); err != nil {
		return nil, err
	}
	if err := validateCustomerID(customerID); err != nil {
		return nil, err
	}

	return &Order{
		id:         id,
		customerID: customerID,
		status:     StatusDraft,
		items:      make(map[ProductID]*OrderItem),
	}, nil
}

func Reconstruct(id OrderID, customerID CustomerID, status Status, items []ItemSnapshot) (*Order, error) {
	if err := validateOrderID(id); err != nil {
		return nil, err
	}
	if err := validateCustomerID(customerID); err != nil {
		return nil, err
	}
	if err := validateStatus(status); err != nil {
		return nil, err
	}
	if status == StatusSubmitted && len(items) == 0 {
		return nil, ErrNoItemsToSubmit
	}

	aggregate := &Order{
		id:         id,
		customerID: customerID,
		status:     status,
		items:      make(map[ProductID]*OrderItem),
	}

	for _, snapshot := range items {
		item, err := newItem(snapshot.ProductID, snapshot.Quantity, snapshot.UnitPrice)
		if err != nil {
			return nil, err
		}
		if _, exists := aggregate.items[item.productID]; exists {
			return nil, ErrDuplicateProduct
		}
		aggregate.items[item.productID] = item
	}

	return aggregate, nil
}

func (o *Order) ID() OrderID {
	return o.id
}

func (o *Order) CustomerID() CustomerID {
	return o.customerID
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) AddItem(productID ProductID, quantity int, unitPrice int) error {
	if err := o.ensureDraft(); err != nil {
		return err
	}
	if err := validateProductID(productID); err != nil {
		return err
	}
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	if unitPrice < 0 {
		return ErrInvalidUnitPrice
	}

	if existing, ok := o.items[productID]; ok {
		if existing.unitPrice != unitPrice {
			return ErrUnitPriceMismatch
		}
		existing.quantity += quantity
		return nil
	}

	item, err := newItem(productID, quantity, unitPrice)
	if err != nil {
		return err
	}
	o.items[productID] = item
	return nil
}

func (o *Order) ChangeQuantity(productID ProductID, quantity int) error {
	if err := o.ensureDraft(); err != nil {
		return err
	}
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	item, ok := o.items[productID]
	if !ok {
		return ErrItemNotFound
	}
	item.quantity = quantity
	return nil
}

func (o *Order) RemoveItem(productID ProductID) error {
	if err := o.ensureDraft(); err != nil {
		return err
	}
	if _, ok := o.items[productID]; !ok {
		return ErrItemNotFound
	}
	delete(o.items, productID)
	return nil
}

func (o *Order) Submit() error {
	if err := o.ensureDraft(); err != nil {
		return err
	}
	if len(o.items) == 0 {
		return ErrNoItemsToSubmit
	}
	o.status = StatusSubmitted
	return nil
}

func (o *Order) TotalPrice() int {
	total := 0
	for _, item := range o.items {
		total += item.Subtotal()
	}
	return total
}

func (o *Order) Items() []OrderItem {
	items := make([]OrderItem, 0, len(o.items))
	for _, item := range o.items {
		items = append(items, *item)
	}
	return items
}

func (o *Order) ensureDraft() error {
	if o.status != StatusDraft {
		return ErrOrderSubmitted
	}
	return nil
}

func validateOrderID(id OrderID) error {
	if strings.TrimSpace(string(id)) == "" {
		return ErrEmptyOrderID
	}
	return nil
}

func validateCustomerID(id CustomerID) error {
	if strings.TrimSpace(string(id)) == "" {
		return ErrEmptyCustomerID
	}
	return nil
}

func validateProductID(id ProductID) error {
	if strings.TrimSpace(string(id)) == "" {
		return ErrEmptyProductID
	}
	return nil
}

func validateStatus(status Status) error {
	switch status {
	case StatusDraft, StatusSubmitted:
		return nil
	default:
		return ErrInvalidStatus
	}
}

func newItem(productID ProductID, quantity int, unitPrice int) (*OrderItem, error) {
	if err := validateProductID(productID); err != nil {
		return nil, err
	}
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}
	if unitPrice < 0 {
		return nil, ErrInvalidUnitPrice
	}

	return &OrderItem{
		productID: productID,
		quantity:  quantity,
		unitPrice: unitPrice,
	}, nil
}
