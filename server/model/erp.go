package model

type ErpDetails struct {
	Data struct {
		EmployeeId string `json:"employeeId"`
		SkypeId    struct {
			Email string `json:"email"`
		} `json:"skypeId"`
	} `json:"data"`
}
