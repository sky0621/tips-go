package fnb

import (
	"github.com/sky0621/tips-go/try/xerrors/after3/fnc"
	"golang.org/x/xerrors"
)

func Exec() error {
	e := fnc.Exec()
	if e != nil {
		return xerrors.Errorf("fnb.Exec error occurred: %w", e)
	}
	return nil
}
