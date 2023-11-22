package auth

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

func LoginService(ctx *gin.Context, emailId string, name string, empId string) {
	var loginDetails model.Login
	var registerDetails model.Register

	// Check if the user already exists in the database
	if db.RecordExist("users", emailId, "email") {
		fmt.Println("LOGIN APIS")
		// Invoke the Login API
		loginReturn := LoginApi(loginDetails, emailId, ctx)
		if !loginReturn {
			fmt.Println("login return is:", loginReturn)
			return
		}
	} else {
		// The user does not exist, proceed with signup process
		fmt.Println("SIGNUP following LOGIN")

		// Invoke the Signup API
		signupReturn := SignupApi(registerDetails, emailId, name, empId, ctx)
		if !signupReturn {
			fmt.Println("signup return is:", signupReturn)
			return
		}

		// Signup successful, now invoke the Login API
		loginReturn := LoginApi(loginDetails, emailId, ctx)
		fmt.Println("login return is", loginReturn)
	}
}

func LoginApi(loginDetails model.Login, emailId string, ctx *gin.Context) bool {
	// Set the email and password in the loginDetails struct
	loginDetails.Email = emailId
	loginDetails.Password = "12345678"

	// Marshal the loginDetails struct into JSON
	loginMarshalData, err := json.Marshal(&loginDetails)
	if err != nil {
		// Handle encoding error and show appropriate response
		response.ShowResponse(
			utils.ENCODING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		return false
	}

	// Create a new HTTP POST request to the Mattermost login URL
	reqst, err := http.NewRequest(utils.REQUEST_POST, utils.MATTERMOST_LOGIN_URL, bytes.NewBuffer(loginMarshalData))
	if err != nil {
		// Handle error while creating the request and show appropriate response
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false
	}

	// Add custom header to the request
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	//fmt.Println("request in login is", reqst.Header)

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		// Handle error while performing the request and show appropriate response
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false
	}

	//fmt.Println("response from login is", resp)

	// Check the response status code
	if resp.StatusCode != 200 {
		// Show appropriate response if the status code is not 200
		response.ShowResponse(utils.ERROR, int64(resp.StatusCode), "", nil, ctx)
		return false
	}

	//fmt.Println("cookies are ", resp.Cookies())
	// Extract the required cookies from the response
	mmauthtoken := resp.Cookies()[0]
	mmuserid := resp.Cookies()[1]
	mmcsrf := resp.Cookies()[2]

	// Create new cookies with the extracted values
	mmauthCookie := &http.Cookie{
		Name:     "MMAUTHTOKEN",
		Value:    mmauthtoken.Value,
		MaxAge:   mmauthtoken.MaxAge,
		Domain:   utils.LOCAL_DOMAIN,
		Path:     "/",
		HttpOnly: false,
	}
	mmuserCookie := &http.Cookie{
		Name:     "MMUSERID",
		Value:    mmuserid.Value,
		MaxAge:   mmuserid.MaxAge,
		Domain:   utils.LOCAL_DOMAIN,
		Path:     "/",
		HttpOnly: false,
	}
	mmcsrfCookie := &http.Cookie{
		Name:     "MMCSRF",
		Value:    mmcsrf.Value,
		MaxAge:   mmcsrf.MaxAge,
		Domain:   utils.LOCAL_DOMAIN,
		Path:     "/",
		HttpOnly: false,
	}

	// Set the cookies in the response writer
	http.SetCookie(ctx.Writer, mmauthCookie)
	http.SetCookie(ctx.Writer, mmuserCookie)
	http.SetCookie(ctx.Writer, mmcsrfCookie)

	// Show the login success response
	// response.ShowResponse(utils.LOGIN_SUCCESSFULL, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

	return true
}

func SignupApi(registerDetails model.Register, emailId string, name string, empId string, ctx *gin.Context) bool {
	fmt.Println("user record not found")

	// Set the email and password in the registerDetails struct
	registerDetails.Email = emailId
	registerDetails.Password = "12345678"

	// Remove spaces from the name and convert it to lowercase
	split := strings.Split(name, " ")
	fullName := ""
	for i := 0; i < len(split); i++ {
		if i != len(split)-1 {

			fullName += split[i] + "_"
		}
		fullName += split[i]
	}

	registerDetails.Username = fullName + "(" + empId + ")"

	fmt.Println("register details:", registerDetails)

	// Marshal the registerDetails struct into JSON
	registerData, err := json.Marshal(&registerDetails)
	if err != nil {
		// Handle encoding error and show appropriate response
		response.ShowResponse(
			utils.ENCODING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		return false
	}

	// Create a new HTTP POST request to the Mattermost signup URL
	reqst, err := http.NewRequest(utils.REQUEST_POST, utils.MATTERMOST_SIGNUP_URL, bytes.NewBuffer(registerData))
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")

	if err != nil {
		// Handle error while creating the request and show appropriate response
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false
	}

	fmt.Println("signup request is ", reqst)
	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		// Handle error while performing the request and show appropriate response
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return false
	}

	fmt.Println("sign up resposne is ", resp)
	// Check the response status code
	if resp.StatusCode != 201 {
		// Show appropriate response if the status code is not 201
		response.ShowResponse(utils.ERROR, int64(resp.StatusCode), "", nil, ctx)
		return false
	}
	fmt.Println("Response from signup:", resp)
	return true
}
