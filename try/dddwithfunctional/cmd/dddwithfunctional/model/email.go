package model

import (
	"errors"
	"strings"
)

var _ ValueObject[string] = (*Email)(nil)

type Email struct {
	value string
}

func (e *Email) Value() string {
	return e.value
}

func NewEmail(value string) (*Email, error) {
	v := strings.TrimSpace(value)
	if len(v) == 0 {
		return nil, errors.New("email must not be empty")
	}
	if !strings.Contains(v, "@") {
		return nil, errors.New("email must contain '@'")
	}
	return &Email{value: v}, nil
}
