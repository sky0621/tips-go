package httpdoc

import "net/http"

type movieHandler struct{}

func (h *movieHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}

type createUserRequest struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Attribute attribute `json:"attribute"`
}

type attribute struct {
	Birthday string `json:"birthday,omitempty"`
	Gender   string `json:"gender,omitempty"`
}
