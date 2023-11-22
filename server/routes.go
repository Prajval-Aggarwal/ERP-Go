package server

import (
	_ "main/docs"

	provider "main/server/gateway"
	"main/server/handler"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(server *Server) {

	//aloowing cors to each route
	server.engine.Use(provider.CORSMiddleware())
	server.engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "success",
		})
	})
	server.engine.GET("/login", provider.UserDetailsMiddleware, handler.LoginHandler)
	server.engine.POST("/logout", provider.UserDetailsMiddleware, handler.LogoutHandler)
	server.engine.POST("/reset", provider.ResetMiddleware, handler.ResetHandler)
	server.engine.GET("/get-cookies", provider.UserDetailsMiddleware, handler.CookieHandler)

	//Channel routes
	server.engine.POST("/channel/create", provider.UserDetailsMiddleware, handler.CreateChannelHandler)
	server.engine.POST("/channel/add-users", provider.UserDetailsMiddleware, handler.AddUsersHandler)
	server.engine.DELETE("/channel/remove-users", provider.UserDetailsMiddleware, handler.RemoveUsersHandler)
	server.engine.POST("/channel/delete-inactive-user", provider.UserDetailsMiddleware, handler.DeleteInactiveUsersHandler)

	server.engine.GET("/get-channels", handler.ChannelHandler)

}
