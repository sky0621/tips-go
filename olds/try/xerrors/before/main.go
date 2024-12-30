package main

import (
	"fmt"

	"github.com/sky0621/tips-go/try/xerrors/before/fna"
)

func main() {
	e := fna.Exec()
	if e != nil {
		lastErr := fmt.Errorf("[before]main error occurred: %w", e)
		fmt.Printf("%+v\n", lastErr)
	}
}
