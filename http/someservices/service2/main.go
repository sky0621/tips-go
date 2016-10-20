package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sv2/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
        <head>
          <title>Service2</title>
        </head>
        <body>
          <h1>Service2</h1>
        </body>
      </html>
      `))
	})
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
