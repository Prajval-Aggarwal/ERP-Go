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
	err := db.QueryExecutor(query, &userId, emailId)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("user id: ", userId)
	var sessionToken string
	query = "SELECT token from sessions where userid=?"

	err = db.QueryExecutor(query, &sessionToken, userId)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	fmt.Println("session token: ", sessionToken)

	//call mattermost logout api
	reqst, err := http.NewRequest(utils.REQUEST_POST, utils.MATTERMOST_LOGOUT_URL, nil)

	if err != nil {
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return

	}
	reqst.Header.Add(utils.CUSTOM_HEADER_KEY_1, utils.CUSTOM_HEADER_VALUE_1)
	reqst.Header.Add("Cookie", "MMAUTHTOKEN="+sessionToken)

	//reqst.Header.Add("Origin", ctx.Request.Header.Get("Origin"))
	fmt.Println("request: ", reqst)
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return
	}
	fmt.Println("response is:", resp)
	if resp.StatusCode != 200 {
		response.ShowResponse(utils.ERROR, int64(resp.StatusCode), "", nil, ctx)
		return
	}
	response.ShowResponse(utils.LOGOUT_SUCCESSFULL, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
