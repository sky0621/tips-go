package model

import (
	"errors"
	"strings"
)

var _ ValueObject[string] = (*UserName)(nil)

type UserName struct {
	value string
}

func (n *UserName) Value() string {
	return n.value
}

func NewUserName(value string) (*UserName, error) {
	v := strings.TrimSpace(value)
	if len(v) == 0 {
		return nil, errors.New("name must not be empty")
	}
	if len(v) > 20 {
		return nil, errors.New("name must be less than or equal to 20 characters")
	}
	return &UserName{value: v}, nil
}
