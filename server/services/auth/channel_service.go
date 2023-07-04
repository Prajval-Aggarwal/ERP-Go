package auth

import (
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"

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
