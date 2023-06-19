package request

type ResetPasswordRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}
