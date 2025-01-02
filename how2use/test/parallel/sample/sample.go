package sample

import (
	"time"
)

func ForParallel() string {
	time.Sleep(3 * time.Second)
	return "done"
}
