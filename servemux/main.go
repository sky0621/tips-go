package main

import (
	"log"
	"net/http"
	"time"

	"github.com/stretchr/graceful"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mux/", withString("testtest", func(w http.ResponseWriter, r *http.Request) {
		log.Println("   func")
	}))
	graceful.Run(":8181", 1*time.Second, mux)
}

// デコレータ―パターンが使える
func withString(str string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[DECORATE]start")
		log.Println(str)
		f(w, r)
		log.Println("[DECORATE]end")
	}
}
