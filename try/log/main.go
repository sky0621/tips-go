package main

import (
	"tips-go/try/log/domain"
	"tips-go/try/log/logger"
)

func main() {
	lgr := logger.NewAppLogger()
	user := domain.NewUser(lgr)
	user.Hello()
}
