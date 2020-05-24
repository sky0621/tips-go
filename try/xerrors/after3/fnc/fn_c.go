package fnc

import (
	"github.com/sky0621/tips-go/try/xerrors/after3/fnd"
	"golang.org/x/xerrors"
)

func Exec() error {
	e := fnd.Exec()
	if e != nil {
		return xerrors.Errorf("fnc.Exec error occurred: %w", e)
	}
	return nil
}
