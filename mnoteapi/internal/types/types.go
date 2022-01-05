// Code generated by goctl. DO NOT EDIT.
package types

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
	Expire   int64  `json:"expire"`
}
