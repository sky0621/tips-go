package adapter

import (
	"context"
	"fmt"

	"github.com/sky0621/tips-go/try/ormtest/src/repository"

	"github.com/jmoiron/sqlx"

	ormtest "github.com/sky0621/tips-go/try/ormtest/src"
)

func NewCustomerService(db *sqlx.DB) ormtest.CustomerService {
	return &customerAdapter{db}
}

type customerAdapter struct {
	db *sqlx.DB
}

func (a *customerAdapter) Customers(ctx context.Context) ([]*ormtest.Customer, error) {
	models, err := repository.Customers().All(ctx, a.db)
	if err != nil {
		return nil, err
	}
	results := []*ormtest.Customer{}
	for _, model := range models {
		results = append(results, &ormtest.Customer{
			ID:       model.ID,
			FullName: fmt.Sprintf("%s %s", model.LastName, model.FirstName),
			Age:      model.Age,
			Nickname: model.Nickname.String,
			Memo:     model.Memo.String,
			IsActive: model.IsActive,
		})
	}
	return results, nil
}
