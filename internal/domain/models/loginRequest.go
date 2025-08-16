package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AppID    int32  `json:"app_id"`
}
