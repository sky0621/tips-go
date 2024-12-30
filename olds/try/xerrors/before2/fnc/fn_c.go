package fnc

import (
	"github.com/pkg/errors"
	"github.com/sky0621/tips-go/try/xerrors/before2/fnd"
)

func Exec() error {
	e := fnd.Exec()
	if e != nil {
		return errors.Wrap(e, "fnc.Exec error occurred")
	}
	return nil
}
