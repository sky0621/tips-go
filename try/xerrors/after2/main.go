package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/after2/util"

	"github.com/sky0621/tips-go/try/xerrors/after2/fna"
)

func main() {
	e := fna.Exec()
	if e != nil {
		lastErr := util.Log(e)
		fmt.Printf("%+v\n", lastErr)
	}
}
