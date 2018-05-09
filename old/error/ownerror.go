package main

import "fmt"

// OwnError ...
type OwnError struct {
	errID  int
	errMsg string
}

// Error ...
func (e OwnError) Error() string {
	return fmt.Sprintf("Own Error - ID:%d Msg:%s\n", e.errID, e.errMsg)
}
