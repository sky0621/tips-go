package fna

import (
	"github.com/sky0621/tips-go/try/xerrors/after2/fnb"
	"github.com/sky0621/tips-go/try/xerrors/after2/util"
)

func Exec() error {
	e := fnb.Exec()
	if e != nil {
		return util.Log(e)
	}
	return nil
}
