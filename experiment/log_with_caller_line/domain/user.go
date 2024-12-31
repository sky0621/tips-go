package domain

import "github.com/sky0621/tips-go/experiment/log_with_caller_line/logger"

type User interface {
	Hello()
}

type user struct {
	lgr logger.AppLogger
}

func NewUser(lgr logger.AppLogger) User {
	return &user{lgr: lgr}
}

func (u *user) Hello() {
	u.lgr.Log("Hello!")
}
