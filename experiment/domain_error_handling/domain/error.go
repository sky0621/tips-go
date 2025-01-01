package domain

import (
	"errors"
	"fmt"
)

type ErrorAttribute struct {
	Key   string
	Value any
}

type Error struct {
	Type       ErrorType
	Message    string           // ログ出力用
	Attributes []ErrorAttribute // ログ出力用
}

func (e *Error) Error() string {
	return fmt.Sprintf("[Type:%s][Message:%s][Attributes:%+v]", e.Type, e.Message, e.Attributes)
}

func (e *Error) Is(target error) bool {
	var err *Error
	ok := errors.As(target, &err)
	return ok
}

type NewDomainError func(attributes ...ErrorAttribute) *Error

func NewError(errorType ErrorType, message string) NewDomainError {
	return func(attributes ...ErrorAttribute) *Error {
		return &Error{Type: errorType, Message: message, Attributes: attributes}
	}
}

var (
	ErrInvalidInput     = NewError(InvalidInput, "invalid input")
	ErrResourceNotFound = NewError(ResourceNotFound, "resource not found")
	ErrUnauthenticated  = NewError(Unauthenticated, "unauthenticated")
	ErrUnauthorized     = NewError(Unauthorized, "unauthorized")
	ErrUnexpectedError  = NewError(UnexpectedError, "unexpected error")
)
