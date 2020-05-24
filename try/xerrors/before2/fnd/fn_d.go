package fnd

import (
	"github.com/pkg/errors"
)

func Exec() error {
	e := errors.New("unexpected error occurred")
	return errors.Wrap(e, "fnd.Exec error occurred")
}
