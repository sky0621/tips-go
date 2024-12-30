package gcfgo

import (
	"fmt"
	"net/http"
)

var count = 0

func CountUp(w http.ResponseWriter, r *http.Request) {
	count++
	if _, err := fmt.Fprintf(w, "Instance execution count: %d", count); err != nil {
		fmt.Println(err)
	}
}
