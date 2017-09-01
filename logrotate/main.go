package main

import (
	"log"

	"time"

	"math/rand"

	"github.com/Sirupsen/logrus"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func main() {
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	logf, err := rotatelogs.New(
		"/data/logs/access_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("/data/logs/access_log"),
		rotatelogs.WithMaxAge(5*time.Minute),
		rotatelogs.WithRotationTime(time.Minute),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(logf)

	for {
		time.Sleep(20 * time.Second)

		logrus.Debugln(rand.Int())
	}

	//al, err := apachelog.New("")
	//if err != nil {
	//	log.Printf("failed to apachelog.New: %s", err)
	//	return
	//}

	//http.ListenAndServe(":8080", al.Wrap(mux, logf))
}
