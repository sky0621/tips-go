package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8181")
		if err != nil {
			fmt.Fprintf(w, "[gowebapp1]ERR:%s", err.Error())
		}
		defer func() {
			if resp == nil {
				return
			}
			if resp.Body == nil {
				return
			}
			resp.Body.Close()
		}()
	})
	http.ListenAndServe(":8080", nil)
}
