package models

type RegRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
