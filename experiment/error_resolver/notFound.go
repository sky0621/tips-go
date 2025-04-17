package main

import (
	"errors"
	"fmt"
)

// sentinel error

var (
	ErrNotFound = errors.New("not found")
)

type NotFoundResolver struct{}

func (r NotFoundResolver) ResolveError(e error) error {
	if errors.Is(e, ErrNotFound) {
		return fmt.Errorf("resource not found: %w", e)
	}
	return e
}
