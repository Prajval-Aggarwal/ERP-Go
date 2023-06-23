package handler

import (
	"main/server/services/auth"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get("emailid")
	name, _ := ctx.Get("name")
	auth.LoginService(ctx, emailId.(string), name.(string))
}
func LogoutHandler(ctx *gin.Context) {
	auth.LogoutService(ctx)
}
func ResetHandler(ctx *gin.Context) {
	password, _ := ctx.Get("password")
	userId, _ := ctx.Get("emailid")
	auth.ResetMattermostPassword(ctx, password.(string), userId.(string))
}
