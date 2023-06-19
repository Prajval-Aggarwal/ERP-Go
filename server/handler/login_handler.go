package handler

import (
	"main/server/services/login"

	"github.com/gin-gonic/gin"
)

func LoginHandler(ctx *gin.Context) {
	emailId, _ := ctx.Get("emailid")
	name, _ := ctx.Get("name")
	login.LoginService(ctx, emailId.(string), name.(string))
}
