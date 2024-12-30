package adapter

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"

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
				// オプション未指定のデフォルトCustomerを生成（１レコード目なのでIDは 1 になるはず）
				if err := CreateCustomer(ctx, db); err != nil {
					t.Fatal(err)
				}

				// 全てのパラメータを上書きしてCustomerを生成
				if err := CreateCustomer(ctx, db,
					withID(2),
					withFirstName("Satoru"),
					withLastName("Sato"),
					withAge(30),
					withNickname("toru"),
					withMemo("メモ")); err != nil {
					t.Fatal(err)
				}
			},
			want: []*ormtest.Customer{
				{
					ID:       1,
					FullName: "ダミー姓 ダミー名",
					Age:      99,
				},
				{
					ID:       2,
					FullName: "Sato Satoru",
					Age:      30,
					Nickname: "toru",
					Memo:     "メモ",
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

type CustomerOption func(*repository.Customer)

func withID(id int64) CustomerOption {
	return func(c *repository.Customer) {
		c.ID = id
	}
}

func withFirstName(firstName string) CustomerOption {
	return func(c *repository.Customer) {
		c.FirstName = firstName
	}
}

func withLastName(lastName string) CustomerOption {
	return func(c *repository.Customer) {
		c.LastName = lastName
	}
}

func withAge(age int) CustomerOption {
	return func(c *repository.Customer) {
		c.Age = age
	}
}

func withNickname(nickname string) CustomerOption {
	return func(c *repository.Customer) {
		c.Nickname = null.StringFrom(nickname)
	}
}

func withMemo(memo string) CustomerOption {
	return func(c *repository.Customer) {
		c.Memo = null.StringFrom(memo)
	}
}

func CreateCustomer(ctx context.Context, db *sqlx.DB, opts ...CustomerOption) error {
	c := repository.Customer{
		FirstName: "ダミー名",
		LastName:  "ダミー姓",
		Age:       99,
	}
	for _, o := range opts {
		o(&c)
	}
	if err := c.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}
