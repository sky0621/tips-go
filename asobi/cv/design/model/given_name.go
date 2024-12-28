package model

import "errors"

var _ ValueObject[string] = (*GivenName)(nil)

type GivenName struct {
	value string
}

func (g *GivenName) GetValue() string {
	return g.value
}

func NewGivenName(value string) (*GivenName, error) {
	if value == "" {
		return nil, errors.New("value is empty")
	}
	return &GivenName{value: value}, nil
}
