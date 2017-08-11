package main

import (
	"fmt"

	"errors"

	"github.com/Sirupsen/logrus"
	"github.com/hashicorp/go-multierror"
)

func main() {
	err := me()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"level":  "error",
		"method": "GET",
		"path":   "/",
		"host":   "example.com",
		"status": 200,
		"bytes":  10234,
		"err":    err,
	}).Error("some error occured")
}

func me() error {
	var result error

	if err := step1(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step2(); err != nil {
		result = multierror.Append(result, err)
	}

	return result
}

func step1() error {
	return errors.New("result is nil.")
}

func step2() error {
	item := "りんご"
	return fmt.Errorf("障害発生：%s", item)
}
