package handler

import (
	"main/server/services/auth"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get(utils.EMAILID)
	name, _ := ctx.Get(utils.NAME)
	token, _ := ctx.Get("token")
	employeeId, _ := ctx.Get("employeeId")
	auth.LoginService(ctx, token.(string), emailId.(string), name.(string), employeeId.(string))
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
