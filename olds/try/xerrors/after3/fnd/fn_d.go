package fnd

import (
	"golang.org/x/xerrors"
)

func Exec() error {
	e := xerrors.New("unexpected error occurred")
	return xerrors.Errorf("fnd.Exec error occurred: %w", e)
}
