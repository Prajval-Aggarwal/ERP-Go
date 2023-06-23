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
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginService(ctx *gin.Context, emailId string, name string) {
	var loginDetails model.Login
	var registerDetails model.Register
	//fmt.Println("Request: ", Request)
	if db.RecordExist("users", emailId, "email") {
		fmt.Println("login api hit only")
		//LOGIN API
		loginReturn := LoginApi(loginDetails, emailId, ctx)
		if !loginReturn {
			fmt.Println("login return is:", loginReturn)
			return
		}

	} else {
		//SIGNUP API
		fmt.Println("signup and login api hit")
		singupReturn := SignupApi(registerDetails, emailId, name, ctx)
		if !singupReturn {
			fmt.Println("sighunup return is:", singupReturn)
			return
		}
		// var userId struct {
		// 	Id string `json:"id"`
		// }
		// query := "SELECT id FROM users WHERE email=?"
		// db.QueryExecutor(query, &userId, emailId)
		// fmt.Println("user id is", userId.Id)

		// var teamId struct {
		// 	Id string `json:"id"`
		// }
		// query = "SELECT id FROM teams where name = 'chicmic1'"
		// db.QueryExecutor(query, &teamId)
		// fmt.Println("teamid: ", teamId.Id)
		// query = "INSERT INTO teammembers(teamid,userid,schemeuser) VALUES(?,?,?)"
		// err := db.RawExecutor(query, teamId.Id, userId.Id, true)
		// if err != nil {
		// 	fmt.Println("error is ", err.Error())
		// 	return
		// }

		loginReturn := LoginApi(loginDetails, emailId, ctx)
		fmt.Println("loginreturn is", loginReturn)

	}

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
	fmt.Println("login details: ", loginDetails)
	reqst, err := http.NewRequest("POST", "http://192.180.0.123:8065/api/v4/users/login", bytes.NewBuffer(loginMarshalData))

	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false

	}
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	//reqst.Header.Add("Origin", ctx.Request.Header.Get("Origin"))
	fmt.Println("request: ", reqst)
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false
	}
	fmt.Println("response is:", resp)
	if resp.StatusCode != 200 {
		response.ShowResponse("Error", int64(resp.StatusCode), "", nil, ctx)
		return false
	}

	fmt.Println("re", resp.Cookies())
	// respCookies := resp.Cookies()
	mmauthtoken := resp.Cookies()[0]
	mmuserid := resp.Cookies()[1]
	mmcsrf := resp.Cookies()[2]

	cookie1 := &http.Cookie{
		Name:     "MMAUTHTOKEN",
		Value:    mmauthtoken.Value,
		MaxAge:   mmauthtoken.MaxAge,
		Path:     "/",
		HttpOnly: false,
	}
	cookie2 := &http.Cookie{
		Name:     "MMUSERID",
		Value:    mmuserid.Value,
		MaxAge:   mmuserid.MaxAge,
		Path:     "/",
		Raw:      mmuserid.Raw,
		HttpOnly: false,
	}

	cookie3 := &http.Cookie{
		Name:     "MMCSRF",
		Value:    mmcsrf.Value,
		MaxAge:   mmcsrf.MaxAge,
		Path:     "/",
		Raw:      mmcsrf.Raw,
		HttpOnly: false,
	}

	// var cookies []*http.Cookie
	// cookies = append(cookies, respCookies...)

	http.SetCookie(ctx.Writer, cookie1)
	http.SetCookie(ctx.Writer, cookie2)
	http.SetCookie(ctx.Writer, cookie3)

	response.ShowResponse("Login  sucessfully", 200, "Success", nil, ctx)

	return true
}

func SignupApi(registerDetails model.Register, emailId string, name string, ctx *gin.Context) bool {
	fmt.Println("record not found")
	registerDetails.Email = emailId
	registerDetails.Password = "123456"

	//remove spce from the name and add _ to it
	lowercase := strings.ToLower(name)
	split := strings.Split(lowercase, " ")
	// joined := strings.Join(split, "_")

	// fmt.Println("name extracted and chaneg is ", name, joined)
	registerDetails.Username = split[0]
	fmt.Println("register details is ", registerDetails)
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
	fmt.Println("resgistered data: ", string(registerData))
	reqst, err := http.NewRequest("POST", "http://192.180.0.123:8065/api/v4/users", bytes.NewBuffer(registerData))
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
	fmt.Println("resposne from signup", resp)
	return true
}
