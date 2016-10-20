package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8181", mux)
}
