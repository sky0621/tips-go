package main

import (
	"errors"
	"fmt"
)

var (
	ErrUnknown = errors.New("unknown error")
)

type OtherResolver struct{}

func (r OtherResolver) ResolveError(e error) error {
	return fmt.Errorf("unexpected error occurred: %w", e)
}
