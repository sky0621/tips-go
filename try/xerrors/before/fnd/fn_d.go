package fnd

import (
	"errors"
	"fmt"
)

func Exec() error {
	e := errors.New("unexpected error occurred")
	return fmt.Errorf("fnd.Exec error occurred: %w", e)
}
