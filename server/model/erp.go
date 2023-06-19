package model

type ErpDetails struct {
	Data struct {
		Name    string `json:"name"`
		SkypeId struct {
			Email string `json:"email"`
		} `json:"skypeId"`
	} `json:"data"`
}
