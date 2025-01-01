package util

import (
	"fmt"
	"time"
)

type Exec func()

func ExecMain(fn Exec) {
	before := time.Now()
	fn()
	fmt.Printf("exec took %v\n", time.Since(before))
}
