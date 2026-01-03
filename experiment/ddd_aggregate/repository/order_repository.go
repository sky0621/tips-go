package repository

import (
	"context"
	"database/sql"

	"github.com/sky0621/tips-go/experiment/ddd_aggregate/infra"
	"github.com/sky0621/tips-go/experiment/ddd_aggregate/order"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Save(ctx context.Context, aggregate *order.Order) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	q := infra.New(tx)
	err = q.UpsertOrder(ctx, infra.UpsertOrderParams{
		ID:          string(aggregate.ID()),
		CustomerID:  string(aggregate.CustomerID()),
		OrderStatus: string(aggregate.Status()),
	})
	if err != nil {
		return err
	}

	err = q.DeleteOrderItemsByOrderID(ctx, string(aggregate.ID()))
	if err != nil {
		return err
	}

	for _, item := range aggregate.Items() {
		err = q.CreateOrderItem(ctx, infra.CreateOrderItemParams{
			OrderID:   string(aggregate.ID()),
			ProductID: string(item.ProductID()),
			Quantity:  int32(item.Quantity()),
			UnitPrice: int32(item.UnitPrice()),
		})
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (r *OrderRepository) FindByID(ctx context.Context, id order.OrderID) (*order.Order, error) {
	q := infra.New(r.db)

	row, err := q.GetOrder(ctx, string(id))
	if err != nil {
		return nil, err
	}

	orderID, err := order.NewOrderID(row.ID)
	if err != nil {
		return nil, err
	}
	customerID, err := order.NewCustomerID(row.CustomerID)
	if err != nil {
		return nil, err
	}
	status, err := order.ParseStatus(row.OrderStatus)
	if err != nil {
		return nil, err
	}

	itemRows, err := q.ListOrderItemsByOrderID(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	snapshots := make([]order.ItemSnapshot, 0, len(itemRows))
	for _, item := range itemRows {
		productID, err := order.NewProductID(item.ProductID)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, order.ItemSnapshot{
			ProductID: productID,
			Quantity:  int(item.Quantity),
			UnitPrice: int(item.UnitPrice),
		})
	}

	return order.Reconstruct(orderID, customerID, status, snapshots)
}
