package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Exposes an HTTP endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
