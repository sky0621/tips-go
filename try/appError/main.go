package main

import "fmt"

func main() {
	e := throwAppError()
	switch t := e.(type) {
	case *AppError:
		fmt.Printf("<app> %s", t.Error())
	case *UserError:
		fmt.Printf("<user> %s", t.Error())
	}

	e2 := throwUserError()
	switch t := e2.(type) {
	case *AppError:
		fmt.Printf("<app> %s", t.Error())
	case *UserError:
		fmt.Printf("<user> %s", t.Error())
	}
}

func throwAppError() error {
	return &AppError{msg: "アプリケーションのエラー"}
}

func throwUserError() error {
	return &UserError{msg: "ユーザのエラー"}
}

type AppError struct {
	msg string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[APP_ERROR] %s", e.msg)
}

type UserError struct {
	msg string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("[USER_ERROR] %s", e.msg)
}