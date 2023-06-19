package handler

import (
	resetpassword "main/server/services/reset-password"

	"github.com/gin-gonic/gin"
)

func ResetHandler(ctx *gin.Context) {
	password, _ := ctx.Get("password")
	userId, _ := ctx.Get("emailid")
	resetpassword.ResetMattermostPassword(ctx, password.(string), userId.(string))
}
