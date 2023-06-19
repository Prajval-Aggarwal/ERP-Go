package response

import "net/http"

type TokenResponse struct {
	MMAUTHTOKEN *http.Cookie `json:"mmauthtoken"`
	MMUSERID    *http.Cookie `json:"mmuserid"`
	MMCSRF      *http.Cookie `json:"mmcsrf"`
	Token       string       `json:"token"`
}
