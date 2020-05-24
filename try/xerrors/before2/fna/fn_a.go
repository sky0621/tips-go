package fna

import (
	"github.com/pkg/errors"
	"github.com/sky0621/tips-go/try/xerrors/before2/fnb"
)

func Exec() error {
	e := fnb.Exec()
	if e != nil {
		return errors.Wrap(e, "fna.Exec error occurred")
	}
	return nil
}
