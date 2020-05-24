package fnd

import (
	"github.com/sky0621/tips-go/try/xerrors/after2/util"
	"golang.org/x/xerrors"
)

func Exec() error {
	e := xerrors.New("unexpected error occurred")
	return util.Log(e)
}
