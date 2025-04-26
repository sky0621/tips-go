package a

import (
	"fmt"
	"github.com/sky0621/tips-go/try/temp3/b"
)

func Aa() error {
	return fmt.Errorf("%w", b.Bb())
}
