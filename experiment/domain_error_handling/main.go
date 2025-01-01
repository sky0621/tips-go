package main

import (
	"log"
	"net/http"

	"github.com/sky0621/tips-go/experiment/domain_error_handling/handler"
)

func main() {
	http.HandleFunc("/api/v1/user", handler.GetUserByIDHandler)

	log.Print("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
