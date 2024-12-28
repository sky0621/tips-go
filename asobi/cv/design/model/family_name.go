package model

import "errors"

var _ ValueObject[string] = (*FamilyName)(nil)

type FamilyName struct {
	value string
}

func (g *FamilyName) GetValue() string {
	return g.value
}

func NewFamilyName(value string) (*FamilyName, error) {
	if value == "" {
		return nil, errors.New("value is empty")
	}
	return &FamilyName{value: value}, nil
}
