package auth

import (
	"encoding/json"
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

type CookieGet struct {
	Mmauthtoken string `json:"auth"`
	Mmcsrftoken string `json:"csrf"`
	Mmuserid    string `json:"userid"`
}

type SessionProps struct {
	OS          string `json:"os"`
	CSRF        string `json:"csrf"`
	IsSaml      string `json:"isSaml"`
	Browser     string `json:"browser"`
	IsMobile    string `json:"isMobile"`
	IsGuest     string `json:"is_guest"`
	Platform    string `json:"platform"`
	IsOAuthUser string `json:"isOAuthUser"`
}

func CookieService(ctx *gin.Context, emailId string) {
	fmt.Println("emailid: ", emailId)
	var userId string
	query := "SELECT id from users where email=?"
	err := db.QueryExecutor(query, &userId, emailId)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("user id: ", userId)
	var cookieGet CookieGet
	cookieGet.Mmuserid = userId
	query = "SELECT token from sessions where userid=?"
	err = db.QueryExecutor(query, &cookieGet.Mmauthtoken, userId)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	query = "SELECT props FROM sessions WHERE userid = '" + userId + "' ORDER BY createat DESC LIMIT 1;"
	var props string
	err = db.QueryExecutor(query, &props)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("", props)
	// fmt.Printf("props: %T", props)
	var temp SessionProps
	err = json.Unmarshal([]byte(props), &temp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("temp is", temp)
	cookieGet.Mmcsrftoken = temp.CSRF
	fmt.Println("cookie: ", cookieGet)

	response.ShowResponse("Cookies are ", utils.HTTP_OK, utils.SUCCESS, cookieGet, ctx)

}
