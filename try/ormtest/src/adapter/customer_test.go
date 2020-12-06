package adapter

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	ormtest "github.com/sky0621/tips-go/try/ormtest/src"
	"github.com/sky0621/tips-go/try/ormtest/src/repository"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		/*
		 * TODO: テストケースは、この順番じゃないと成功しない。本当は"no records"と"some records"はテストケースを分けるべき。
		 */
		{
			name:        "no records",
			prepareFunc: func() {},
			want:        []*ormtest.Customer{},
			wantError:   false,
		},
		{
			name: "some records",
			prepareFunc: func() {
				c1 := repository.Customer{
					ID:        1,
					FirstName: "Satoru",
					LastName:  "Sato",
					Age:       30,
					Nickname:  null.StringFrom("toru"),
					Memo:      null.StringFrom("上客"),
					IsActive:  false,
				}
				if err := c1.Insert(ctx, db, boil.Greylist(repository.CustomerColumns.IsActive)); err != nil {
					t.Fatal(err)
				}

				c2 := repository.Customer{
					ID:        2,
					FirstName: "Takeru",
					LastName:  "Tanaka",
					Age:       19,
				}
				if err := c2.Insert(ctx, db, boil.Infer()); err != nil {
					t.Fatal(err)
				}
			},
			want: []*ormtest.Customer{
				{
					ID:       1,
					FullName: "Sato Satoru",
					Age:      30,
					Nickname: "toru",
					Memo:     "上客",
					IsActive: false,
				},
				{
					ID:       2,
					FullName: "Tanaka Takeru",
					Age:      19,
					IsActive: true,
				},
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareFunc()

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
