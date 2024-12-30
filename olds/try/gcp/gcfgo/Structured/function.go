package Structured

import (
	"Structured/subpkg"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

var lgr *zap.Logger

func init() {
	var err error
	lgr, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func StructuredHello(w http.ResponseWriter, r *http.Request) {
	var jf subpkg.JsonForm
	if err := json.NewDecoder(r.Body).Decode(&jf); err != nil {
		lgr.Error(err.Error())
	}
	if _, err := fmt.Fprintf(w, "Hello, %s(%s)", jf.Name, jf.ID); err != nil {
		lgr.Error(err.Error())
	}
}
