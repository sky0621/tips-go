//go:generate stringer -type Signal signal.go
package main

type Signal int

const (
	Green Signal = iota + 1
	Yellow
	Red
)
