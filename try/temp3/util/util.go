package util

import (
	"fmt"
	"runtime"
)

func FileWithLineNum() string {
	_, name, line, ok := runtime.Caller(2)
	if !ok {
		return "-"
	}
	return fmt.Sprintf("%s:%d", name, line)
}
