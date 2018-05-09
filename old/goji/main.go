package main

import (
	"fmt"
	"net/http"

	goji "goji.io"

	"goji.io/pat"

	"golang.org/x/net/context"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hi/:name"), hi)
	http.ListenAndServe(":7171", mux)
}

func hi(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hi %s", name)
}
