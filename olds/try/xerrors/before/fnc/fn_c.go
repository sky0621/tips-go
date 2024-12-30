package fnc

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/before/fnd"
)

func Exec() error {
	e := fnd.Exec()
	if e != nil {
		return fmt.Errorf("fnc.Exec error occurred: %w", e)
	}
	return nil
}
