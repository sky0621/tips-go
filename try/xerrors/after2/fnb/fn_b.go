package fnb

import (
	"github.com/sky0621/tips-go/try/xerrors/after2/fnc"
	"github.com/sky0621/tips-go/try/xerrors/after2/util"
)

func Exec() error {
	e := fnc.Exec()
	if e != nil {
		return util.Log(e)
	}
	return nil
}
