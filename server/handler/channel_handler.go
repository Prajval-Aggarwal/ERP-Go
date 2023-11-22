package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/auth"
	"main/server/services/channel"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func ChannelHandler(ctx *gin.Context) {
	userId := ctx.Query("userId")

	auth.GetChannelService(ctx, userId)
}

func DeleteInactiveUsersHandler(ctx *gin.Context) {

	var req request.DeleteInactiveMemeber
	err := utils.RequestDecoding(ctx, &req)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	channel.DeleteInactiveMembers(ctx, &req)

}
