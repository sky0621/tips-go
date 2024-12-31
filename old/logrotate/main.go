package main

import (
	"log"
	"os"

	"time"

	"math/rand"

	"io"

	"github.com/Sirupsen/logrus"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	logf, err := rotatelogs.New(
		"./access_log.%Y%m%d%H%M%S",
		rotatelogs.WithLinkName("./access_log"),
		rotatelogs.WithMaxAge(10*time.Minute),
		rotatelogs.WithRotationTime(10*time.Second),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	out := io.MultiWriter(os.Stdout, logf)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(out)

	for {
		time.Sleep(3 * time.Second)
		//fmt.Println("logrus.Debugln")

		logrus.Debugln(rand.Int())
	}

	//al, err := apachelog.New("")
	//if err != nil {
	//	log_with_caller_line.Printf("failed to apachelog.New: %s", err)
	//	return
	//}

	//http.ListenAndServe(":8080", al.Wrap(mux, logf))
}
