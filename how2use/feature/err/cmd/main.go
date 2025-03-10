package main

import (
	"errors"
	"fmt"
	"github.com/sky0621/tips-go/how2use/feature/err"
)

func main() {
	e := fn(100)

	switch t := e.(type) {
	case *err.AppError:
		fmt.Println(t.Error())
	default:
		fmt.Println("EEEEE")
	}

	if e != nil {
		var appErr *err.AppError
		if errors.As(e, &appErr) {
			fmt.Println(e.Error())
		} else {
			fmt.Println("OOOOO")
		}
	}

	e2 := fn(200)

	if e2 != nil {
		var appErr *err.AppError2
		if errors.As(e2, &appErr) {
			fmt.Println(e2.Error())
		} else {
			fmt.Println("RRRRR")
		}
	}
}

func fn(code int) error {
	if code == 100 {
		return err.NewAppError(code, errors.New("app error"))
	}
	if code == 200 {
		return err.NewAppError2(code, "ERR200", errors.New("app error"))
	}
	return nil
}
