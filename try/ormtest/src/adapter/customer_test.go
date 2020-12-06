package adapter

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	ormtest "github.com/sky0621/tips-go/try/ormtest/src"
)

func TestCustomers(t *testing.T) {
	db := setupDB()
	defer teardownDB(db)

	service := NewCustomerService(db)
	ctx := context.Background()

	tests := []struct {
		name        string
		prepareFunc func()
		want        []*ormtest.Customer
		wantError   bool
	}{
		{
			name:        "no records",
			prepareFunc: func() {},
			want:        []*ormtest.Customer{},
			wantError:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Customers(ctx)
			if (err != nil) != tt.wantError {
				t.Errorf("error = %v, wantError = %v", err, tt.wantError)
				return
			}
			opts := []cmp.Option{}
			if diff := cmp.Diff(tt.want, got, opts...); diff != "" {
				t.Errorf("unmatch (-want +got):\n%s", diff)
				return
			}
		})
	}
}
