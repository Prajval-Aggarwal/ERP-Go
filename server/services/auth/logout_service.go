package auth

import (
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutService(ctx *gin.Context, emailId string) {

	fmt.Println("email id:", emailId)
	var userId string
	query := "SELECT id from users where email=?"
	db.QueryExecutor(query, &userId, emailId)
	fmt.Println("user id: ", userId)
	var sessionToken string
	query = "SELECT token from sessions where userid=?"

	db.QueryExecutor(query, &sessionToken, userId)

	fmt.Println("session token: ", sessionToken)

	// cookies, _ := ctx.Request.Cookie("MMAUTHTOKEN")
	// fmt.Println("cookies", cookies.Value)
	// //deleting the cookies
	// mmauthCookie := &http.Cookie{
	// 	Name:     "MMAUTHTOKEN",
	// 	Value:    sessionToken,
	// 	MaxAge:   0,
	// 	Path:     "/",
	// 	HttpOnly: false,
	// }

	//call mattermost logout api
	reqst, err := http.NewRequest("POST", "http://192.180.0.123:8065/api/v4/users/logout", nil)

	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return

	}
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	reqst.Header.Add("Cookie", "MMAUTHTOKEN="+sessionToken)
	// http.SetCookie(ctx.Writer, mmauthCookie)
	//reqst.Header.Add("Origin", ctx.Request.Header.Get("Origin"))
	fmt.Println("request: ", reqst)
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return
	}
	fmt.Println("response is:", resp)
	if resp.StatusCode != 200 {
		response.ShowResponse("Error", int64(resp.StatusCode), "", nil, ctx)
		return
	}
	response.ShowResponse("Logout sucessfully", 200, "Success", nil, ctx)
}
