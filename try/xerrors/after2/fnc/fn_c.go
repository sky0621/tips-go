package fnc

import (
	"github.com/sky0621/tips-go/try/xerrors/after2/fnd"
	"github.com/sky0621/tips-go/try/xerrors/after2/util"
)

func Exec() error {
	e := fnd.Exec()
	if e != nil {
		return util.Log(e)
	}
	return nil
}
