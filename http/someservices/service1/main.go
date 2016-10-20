package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sv1/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <title>Service1</title>
        </head>
        <body>
          <h1>Service1</h1>
        </body>
      </html>
      `))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
