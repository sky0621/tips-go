package b

import (
	"fmt"
	"github.com/sky0621/tips-go/try/temp3/c"
)

func Bb() error {
	return fmt.Errorf("%w", c.Cc())
}
