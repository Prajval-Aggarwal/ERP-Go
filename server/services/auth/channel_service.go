package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetChannelService(ctx *gin.Context, userId string) {
	var data []string
	query := "SELECT  channelid FROM channelmembers WHERE userid = ?"
	err := db.QueryExecutor(query, &data, userId)
	if err != nil {
		fmt.Println("error is", err)
		return
	}
	fmt.Println("channels data is", data)

	response.ShowResponse(utils.DATA_FETCH_SUCESS, utils.HTTP_OK, utils.SUCCESS, data, ctx)
}

func CreateChannel(ctx *gin.Context, emailId string, req request.CreateChannelRequest) {
	var loginDetails model.Login
	if LoginApi(loginDetails, emailId, ctx) {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
		return
	}

	var token string
	query := "SELECT s.token FROM sessions s JOIN users u ON s.userid=u.id WHERE u.email=? order by s.createat DESC LIMIT 1"
	err := db.QueryExecutor(query, &token, emailId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	fmt.Println("Token is", token)

	var teamId string
	query = "SELECT id FROM teams WHERE name='chicmic'"
	err = db.QueryExecutor(query, &teamId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var bearer = "Bearer " + token

	fmt.Println("Team Id is", teamId)

	chanelCreationData := struct {
		TeamId      string `json:"team_id"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
	}{
		TeamId:      teamId,
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Type:        req.Type,
	}

	jsonFormat, err := json.Marshal(&chanelCreationData)
	if err != nil {
		// Handle encoding error and show appropriate response
		return
	}
	reqst, err := http.NewRequest(utils.REQUEST_POST, utils.CHANNEL_URL, bytes.NewBuffer(jsonFormat))
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//Adding required headers
	reqst.Header.Add("Authorization", bearer)

	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

		return
	}

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	//resp.Body.Close()

	// print response body
	fmt.Println("hsdavsd", string(body))
	var res interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		// response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		fmt.Println("error", err)
		return
	}
	fmt.Println("res------->", res)

	if resp.StatusCode != 201 {
		// Show appropriate response if the status code is not 201
		response.ShowResponse(utils.ERROR, int64(resp.StatusCode), utils.FAILURE, res, ctx)
		return
	}

	response.ShowResponse("Channel created successfully", utils.HTTP_OK, utils.SUCCESS, res, ctx)

}

func AddUsersToChannel(ctx *gin.Context, emailId string, req request.AddMemeber) {
	var token string
	query := "SELECT s.token FROM sessions s JOIN users u ON s.userid=u.id WHERE u.email=? order by s.createat DESC LIMIT 1"
	err := db.QueryExecutor(query, &token, emailId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	fmt.Println("token is", token)
	for _, email := range req.EmailIds {
		var id string
		query := "SELECT id from users WHERE email=?"
		err := db.QueryExecutor(query, &id, email)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		reqstBody := struct {
			UserId string `json:"user_id"`
		}{
			UserId: id,
		}
		jsonFormat, err := json.Marshal(&reqstBody)
		if err != nil {
			fmt.Println("askhdjhadjsajdbj")
			return
		}

		reqst, err := http.NewRequest("POST", utils.CHANNEL_URL+"/"+req.ChannelId+"/members", bytes.NewBuffer(jsonFormat))
		//Correct it
		var bearer = "Bearer " + token
		reqst.Header.Add("Authorization", bearer)

		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		resp, err := http.DefaultClient.Do(reqst)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

			return
		}

		body, error := ioutil.ReadAll(resp.Body)
		if error != nil {
			fmt.Println(error)
		}
		// close response body
		//resp.Body.Close()

		// print response body
		var res interface{}
		err = json.Unmarshal(body, &res)
		if err != nil {
			// response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			fmt.Println("error", err)
			return
		}
		fmt.Println("hsdavsd", string(body))
		if resp.StatusCode != 201 {
			// Show appropriate response if the status code is not 201
			response.ShowResponse(utils.ERROR, int64(resp.StatusCode), utils.FAILURE, res, ctx)
			return
		}

	}
	response.ShowResponse("Users added to channel sucessfully", utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func RemoveUsersFromChannel(ctx *gin.Context, emailId string, req request.AddMemeber) {

	var token string
	query := "SELECT s.token FROM sessions s JOIN users u ON s.userid=u.id WHERE u.email=? order by s.createat DESC LIMIT 1"
	err := db.QueryExecutor(query, &token, emailId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	for _, email := range req.EmailIds {
		var id string
		query := "SELECT id from users WHERE email=?"
		err := db.QueryExecutor(query, &id, email)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		reqst, err := http.NewRequest("DELETE", utils.CHANNEL_URL+"/"+req.ChannelId+"/members/"+id, nil)
		//Correct it
		var bearer = "Bearer " + token
		reqst.Header.Add("Authorization", bearer)

		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		resp, err := http.DefaultClient.Do(reqst)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

			return
		}

		fmt.Println("resposfjdanflksadf", resp)
		if resp.StatusCode != 200 {
			// Show appropriate response if the status code is not 201
			response.ShowResponse(utils.ERROR, int64(resp.StatusCode), "", nil, ctx)
			return
		}

	}

	response.ShowResponse("Users removed from channel sucessfully", utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
