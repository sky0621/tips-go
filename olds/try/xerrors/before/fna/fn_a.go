package fna

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/before/fnb"
)

func Exec() error {
	e := fnb.Exec()
	if e != nil {
		return fmt.Errorf("fna.Exec error occurred: %w", e)
	}
	return nil
}
