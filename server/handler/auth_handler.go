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
	emailId, _ := ctx.Get("emailid")
	auth.LogoutService(ctx, emailId.(string))
}
func ResetHandler(ctx *gin.Context) {
	password, _ := ctx.Get("password")
	userId, _ := ctx.Get("emailid")
	auth.ResetMattermostPassword(ctx, password.(string), userId.(string))
}

func CookieHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get("emailid")
	auth.CookieService(ctx, emailId.(string))
}
