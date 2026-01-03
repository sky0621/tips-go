package order

type OrderItem struct {
	productID ProductID
	quantity  int
	unitPrice int
}

func (i OrderItem) ProductID() ProductID {
	return i.productID
}

func (i OrderItem) Quantity() int {
	return i.quantity
}

func (i OrderItem) UnitPrice() int {
	return i.unitPrice
}

func (i OrderItem) Subtotal() int {
	return i.quantity * i.unitPrice
}
