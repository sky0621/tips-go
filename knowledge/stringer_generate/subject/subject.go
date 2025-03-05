//go:generate go run golang.org/x/tools/cmd/stringer -type=Subject subject.go
package subject

type Subject int

const (
	Japanese Subject = iota + 1
	Math
	English
	Science
	Social
)
