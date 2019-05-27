package domain

import "tips-go/try/log/logger"

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
