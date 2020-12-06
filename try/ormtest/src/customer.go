package ormtest

import (
	"context"
)

type CustomerService interface {
	Customers(context.Context) ([]*Customer, error)
}

type Customer struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullName"`
	Age      int    `json:"age"`
	Nickname string `json:"nickname,omitempty"`
	Memo     string `json:"memo,omitempty"`
	IsActive bool   `json:"isActive"`
}
