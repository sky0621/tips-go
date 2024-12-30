package fna

import (
	"github.com/sky0621/tips-go/try/xerrors/after/fnb"
	"golang.org/x/xerrors"
)

func Exec() error {
	e := fnb.Exec()
	if e != nil {
		return xerrors.Errorf("fna.Exec error occurred: %w", e)
	}
	return nil
}
