package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/after3/fna"
	"golang.org/x/xerrors"
)

func main() {
	e := fna.Exec()
	if e != nil {
		lastErr := xerrors.Errorf("[after3]main error occurred: %w", e)
		fmt.Printf("%+v\n", xerrors.Unwrap(lastErr))
	}
}
