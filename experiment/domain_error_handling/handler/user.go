package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sky0621/tips-go/experiment/domain_error_handling/domain"
)

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	strID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := domain.NewUserService().GetUserByID(context.Background(), id)
	if err != nil {
		httpError := ToHTTPError(err)
		log.Printf("Error: %+v\n", httpError)
		http.Error(w, httpError.Error(), httpError.StatusCode.Code())
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, user)
	if err != nil {
		log.Println(err)
		return
	}
}
