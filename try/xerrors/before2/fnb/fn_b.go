package fnb

import (
	"github.com/pkg/errors"
	"github.com/sky0621/tips-go/try/xerrors/before2/fnc"
)

func Exec() error {
	e := fnc.Exec()
	if e != nil {
		return errors.Wrap(e, "fnb.Exec error occurred")
	}
	return nil
}
