package main

import (
	"errors"
	"fmt"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type PermissionDeniedResolver struct{}

func (r PermissionDeniedResolver) ResolveError(e error) error {
	if errors.Is(e, ErrPermissionDenied) {
		return fmt.Errorf("permission denied: %w", e)
	}
	return e
}
