package main

import (
	"fmt"
	"os"

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
	logrus.Debug("This is a debug log_with_caller_line.")
	logrus.Info("This is a info log_with_caller_line.")
	logrus.Warn("This is a warn log_with_caller_line.")
	logrus.Error("This is a error log_with_caller_line.")

	fmt.Println("　　　　　　　　　　　　　　")
	fmt.Println("ログ出力先をファイルに変更")
	logfile, err := os.Create("output.log_with_caller_line")
	if err != nil {
		fmt.Println(err)
	}
	defer logfile.Close()
	logrus.SetOutput(logfile)
	logrus.Warn("Output file.")
}
