package auth

import (
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Value string
}

func GetChannelService(ctx *gin.Context, userId string) {
	var data []string
	query := "select  channelid from channelmembers where userid = ?"
	err := db.QueryExecutor(query, &data, userId)
	if err != nil {
		fmt.Println("error is", err)
		return
	}
	fmt.Println("channels data is", data)

	response.ShowResponse("Fetched sucessfully", utils.HTTP_OK, "success", data, ctx)
}
