package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	fmt.Println("logrusでは、任意のフィールドを定義してログ出力するようにできる")

	logrus.WithFields(logrus.Fields{
		"host": "example.com",
		"path": "/",
		"feel": "happy",
	}).Info("some info.")

	lgrs := logrus.WithFields(logrus.Fields{
		"host": "example.com",
		"path": "/err",
		"feel": "sad",
	})
	lgrs.Error("EEEEERRRRRRoooooorrr")

	fmt.Println("　　　　　　　　　　　　　　")
	fmt.Println("ログレベルを WARN にセットすると、DEBUG や INFO は出なくなる")
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Debug("This is a debug log.")
	logrus.Info("This is a info log.")
	logrus.Warn("This is a warn log.")
	logrus.Error("This is a error log.")
}
