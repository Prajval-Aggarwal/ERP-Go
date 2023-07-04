package handler

import (
	"main/server/services/auth"

	"github.com/gin-gonic/gin"
)

func ChannelHandler(ctx *gin.Context) {
	userId := ctx.Query("userId")

	auth.GetChannelService(ctx, userId)
}
