package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/sky0621/tips-go/experiment/domain_error_handling/domain"
)

type HttpStatusCode int

func (h HttpStatusCode) Code() int {
	return int(h)
}

const (
	StatusOK        HttpStatusCode = http.StatusOK
	StatusCreated   HttpStatusCode = http.StatusCreated
	StatusNoContent HttpStatusCode = http.StatusNoContent

	StatusBadRequest   HttpStatusCode = http.StatusBadRequest
	StatusUnauthorized HttpStatusCode = http.StatusUnauthorized
	StatusForbidden    HttpStatusCode = http.StatusForbidden
	StatusNotFound     HttpStatusCode = http.StatusNotFound

	StatusInternalServerError HttpStatusCode = http.StatusInternalServerError
)

type HTTPError struct {
	StatusCode    HttpStatusCode
	StatusMessage string
}

func (e *HTTPError) Error() string {
	return e.StatusMessage
}

func NewHTTPError(statusCode HttpStatusCode) *HTTPError {
	return &HTTPError{statusCode, http.StatusText(int(statusCode))}
}

func ToHTTPError(err error) *HTTPError {
	var domainError *domain.Error
	if errors.As(err, &domainError) {
		switch domainError.Type {
		case domain.InvalidInput:
			return NewHTTPError(StatusBadRequest)
		case domain.Unauthenticated:
			return NewHTTPError(StatusUnauthorized)
		case domain.Unauthorized:
			return NewHTTPError(StatusForbidden)
		case domain.ResourceNotFound:
			return NewHTTPError(StatusNotFound)
		case domain.UnexpectedError:
			return NewHTTPError(StatusInternalServerError)
		default:
			return NewHTTPError(StatusInternalServerError)
		}
	}
	log.Printf("[OTHER ERROR] %+v", err)
	return NewHTTPError(StatusInternalServerError)
}
