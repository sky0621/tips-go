package main

import (
	"github.com/sky0621/tips-go/experiment/log_with_caller_line/domain"
	"github.com/sky0621/tips-go/experiment/log_with_caller_line/logger"
)

func main() {
	lgr := logger.NewAppLogger()
	user := domain.NewUser(lgr)
	user.Hello()
}
