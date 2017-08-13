package httpdoc2

import "net/http"

type userHandler struct {
}

func (h *userHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// 未実装の状態でもよい
}

type UserRequest struct {
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"`
}

type UserResponse struct {
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"`
}
