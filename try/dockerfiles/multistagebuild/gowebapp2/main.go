package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ba, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, "SERVICE2 <- %s", string(ba))
	})
	http.ListenAndServe(":8181", nil)
}
