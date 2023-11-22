package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/auth"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)
	name, _ := ctx.Get(utils.NAME)
	employeeId, _ := ctx.Get("employeeId")
	auth.LoginService(ctx, emailId.(string), name.(string), employeeId.(string))
}
func LogoutHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)
	auth.LogoutService(ctx, emailId.(string))
}
func ResetHandler(ctx *gin.Context) {
	password, _ := ctx.Get(utils.PASSWORD)
	userId, _ := ctx.Get(utils.EMAILID)
	auth.ResetMattermostPassword(ctx, password.(string), userId.(string))
}

func CookieHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)
	auth.CookieService(ctx, emailId.(string))
}

func CreateChannelHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)

	var req request.CreateChannelRequest
	err := utils.RequestDecoding(ctx, &req)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	auth.CreateChannel(ctx, emailId.(string), req)

}

func AddUsersHandler(ctx *gin.Context) {

	emailId, _ := ctx.Get(utils.EMAILID)
	var req request.AddMemeber
	err := utils.RequestDecoding(ctx, &req)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	auth.AddUsersToChannel(ctx, emailId.(string), req)
}

func RemoveUsersHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)
	var req request.AddMemeber
	err := utils.RequestDecoding(ctx, &req)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	auth.RemoveUsersFromChannel(ctx, emailId.(string), req)
}
