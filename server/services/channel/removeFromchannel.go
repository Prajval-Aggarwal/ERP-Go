package channel

import (
	"fmt"
	"io"
	"log"
	"main/server/db"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func DeleteInactiveMembers(ctx *gin.Context, req *request.DeleteInactiveMemeber) {

	userEmail := req.UserEmail
	//get the user id fromthe db corresponding to the userEmail

	query := `SELECT id FROM users WHERE email=?`
	var userId string

	err := db.QueryExecutor(query, &userId, userEmail)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//get all the channels in which user is present ,then call the leave channel function for each channel

	var channels []string

	query = `SELECT channelid FROM channelmembers WHERE userid=?`

	err = db.QueryExecutor(query, &channels, userId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	// fmt.Println("channel ids----> :", channels)

	//get the token of the user from the session table in db
	token, ok := ctx.Get("token")
	if !ok {
		log.Fatal(err)
	}

	fmt.Println("token is", token.(string))

	//create a request for hit to leavechannel function in mattermost
	//now create request

	// Create a map or struct to hold your parameters

	for _, channel := range channels {

		params := url.Values{}
		params.Set("channel_id", channel)
		params.Set("user_id", userId)

		// Create a new URL by appending query parameters
		urlWithParams := utils.CHANNEL_URL + "/" + channel + "/" + "members" + "/" + userId
		fmt.Println("url--------->", urlWithParams)

		reqst, err := http.NewRequest("DELETE", urlWithParams, nil)
		//Correct it
		var bearer = "Bearer " + token.(string)
		reqst.Header.Add("Authorization", bearer)

		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}

		// fmt.Println("Request for LEAVE CHANNEL-------->", reqst)
		resp, err := http.DefaultClient.Do(reqst)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

			return
		}
		respBytes, _ := io.ReadAll(resp.Body)
		// fmt.Println("response from LEAVE CHANNEL API-------->", resp)
		if resp.StatusCode != 200 && resp.StatusCode != 400 {

			response.ShowResponse(string(respBytes), int64(resp.StatusCode), utils.FAILURE, nil, ctx)
			return
		}

	}

	//Deactivate user from mattermost(hit delete user api of mattermost with userid)

	urlWithParams := utils.MATTERMOST_SIGNUP_URL + "/" + userId
	fmt.Println("url--------->", urlWithParams)

	reqst, err := http.NewRequest("DELETE", urlWithParams, nil)
	//Correct it
	var bearer = "Bearer " + token.(string)
	reqst.Header.Add("Authorization", bearer)

	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// fmt.Println("Request for DEACTIVATE USER-------->", reqst)
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

		return
	}
	respBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println("response from DEACTIVATE USER-------->", resp)

	if resp.StatusCode != 200 {
		response.ShowResponse(string(respBytes), int64(resp.StatusCode), utils.FAILURE, nil, ctx)

	}

	response.ShowResponse("User removed from all Channels", utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}
