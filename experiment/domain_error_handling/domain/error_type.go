//go:generate stringer -type ErrorType error_type.go
package domain

type ErrorType int

const (
	InvalidInput ErrorType = iota + 1
	ResourceNotFound
	Unauthenticated
	Unauthorized
	UnexpectedError
)
