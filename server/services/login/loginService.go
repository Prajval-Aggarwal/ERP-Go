package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginService(ctx *gin.Context, emailId string, name string) {
	var loginDetails model.Login
	var registerDetails model.Register
	//fmt.Println("Request: ", Request)
	if db.RecordExist("users", emailId, "email") {
		//LOGIN API
		loginReturn := LoginApi(loginDetails, emailId, ctx)
		if !loginReturn {
			return
		}

	} else {
		//SIGNUP API
		singupReturn := SignupApi(registerDetails, emailId, name, ctx)
		if !singupReturn {
			return
		}
		loginReturn := LoginApi(loginDetails, emailId, ctx)
		if !loginReturn {
			return
		}
	}

}

func SignupApi(registerDetails model.Register, emailId string, name string, ctx *gin.Context) bool {
	fmt.Println("record not found")
	registerDetails.Email = emailId
	registerDetails.Password = "123456"

	//remove spce from the name and add _ to it
	// lowercase := strings.ToLower(name)
	// split := strings.Split(lowercase, " ")
	// joined := strings.Join(split, "_")

	// fmt.Println("name extracted and chaneg is ", name, joined)
	registerDetails.Username = name

	registerData, err := json.Marshal(&registerDetails)
	if err != nil {
		response.ShowResponse(
			"Error in Encoding data",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		return false
	}
	reqst, err := http.NewRequest("POST", "http://192.180.4.118:8065/api/v4/users", bytes.NewBuffer(registerData))
	//fmt.Println("request: ", reqst)
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		return false

	}
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		return false
	}
	fmt.Println("resposne is", resp)
	return true
}

func LoginApi(loginDetails model.Login, emailId string, ctx *gin.Context) bool {
	//ctx.Header("Referrer-Policy", "no-referrer")
	fmt.Println("record found")
	loginDetails.Email = emailId
	loginDetails.Password = "123456"

	loginMarshalData, err := json.Marshal(&loginDetails)
	if err != nil {
		response.ShowResponse(
			"Error in Encoding data",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		return false
	}

	reqst, err := http.NewRequest("POST", "http://192.180.4.118:8065/api/v4/users/login", bytes.NewBuffer(loginMarshalData))
	fmt.Println("request: ", reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		return false

	}
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		return false
	}

	fmt.Println("re", resp.Cookies())

	mmauthtoken := resp.Cookies()[0]
	mmuserid := resp.Cookies()[1]

	mmcsrf := resp.Cookies()[2]
	// token := resp.Header["Token"][0]
	// tokenResponse := response.TokenResponse{
	// 	MMAUTHTOKEN: mmauthtoken,
	// 	MMUSERID:    mmuserid,
	// 	MMCSRF:      mmcsrf,
	// 	Token:       token,
	// }

	raw := mmauthtoken.Raw
	raw += "; Domain=ngrok-free.app"

	cookie1 := &http.Cookie{
		Name:     "MMAUTHTOKEN",
		Value:    mmauthtoken.Value,
		MaxAge:   mmauthtoken.MaxAge,
		Path:     "/",
		Raw:      raw,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	cookie2 := &http.Cookie{
		Name:     "MMUSERID",
		Value:    mmuserid.Value,
		MaxAge:   mmuserid.MaxAge,
		Path:     "/",
		Raw:      mmuserid.Raw,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	cookie3 := &http.Cookie{
		Name:     "MMCSRF",
		Value:    mmcsrf.Value,
		MaxAge:   mmcsrf.MaxAge,
		Path:     "/",
		Raw:      mmcsrf.Raw,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	var cookies []*http.Cookie
	cookies = append(cookies, cookie1)

	http.SetCookie(ctx.Writer, cookie1)
	http.SetCookie(ctx.Writer, cookie2)
	http.SetCookie(ctx.Writer, cookie3)

	response.ShowResponse("Login  sucessfully", 200, "Success", cookies, ctx)

	return true
}
