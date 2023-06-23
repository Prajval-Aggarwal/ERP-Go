package server

import (
	_ "main/docs"

	provider "main/server/gateway"
	"main/server/handler"
)

func ConfigureRoutes(server *Server) {

	//aloowing cors to each route
	server.engine.Use(provider.CORSMiddleware())
	server.engine.GET("/login", provider.UserDetailsMiddleware, handler.LoginHandler)
	server.engine.DELETE("/logout", handler.LogoutHandler)
	server.engine.POST("/reset", provider.ResetMiddleware, handler.ResetHandler)

}
