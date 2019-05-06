package Quickstart

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "Hello, World! (go111)"); err != nil {
		fmt.Println(err)
	}
}
