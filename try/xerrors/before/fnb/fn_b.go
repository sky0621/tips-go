package fnb

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/before/fnc"
)

func Exec() error {
	e := fnc.Exec()
	if e != nil {
		return fmt.Errorf("fnb.Exec error occurred: %w", e)
	}
	return nil
}
