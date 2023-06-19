package model

type Login struct {
	Email    string `json:"login_id"`
	Password string `json:"password"`
}
