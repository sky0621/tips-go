package err

import "fmt"

type AppError struct {
	code int
	err  error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code:%d err:%v", e.code, e.err)
}

func NewAppError(code int, err error) *AppError {
	return &AppError{code, err}
}

type AppError2 struct {
	code int
	kind string
	err  error
}

func (e *AppError2) Error() string {
	return fmt.Sprintf("code:%d kind:%s err:%v", e.code, e.kind, e.err)
}

func NewAppError2(code int, kind string, err error) *AppError2 {
	return &AppError2{code, kind, err}
}
