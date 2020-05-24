package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sky0621/tips-go/try/xerrors/before2/fna"
)

func main() {
	e := fna.Exec()
	if e != nil {
		lastErr := errors.Wrap(e, "[before2]main error occurred")
		fmt.Printf("%+v\n", errors.Unwrap(lastErr))
	}
}
