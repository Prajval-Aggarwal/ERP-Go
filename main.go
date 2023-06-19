package main

import (
	"log"
	"main/server"
	"main/server/db"
	"main/server/socket"
	"os"

	"github.com/joho/godotenv"
)

// @title Gin Demo App
// @version 1.0
// @description This is a demo version of Gin app.
// @BasePath /
func main() {
	if err := godotenv.Load("/home/chicmic/go/src/github.com/Prajval-Aggarwal/mattermost-go-server/.env"); err != nil {
		log.Fatal("Error loadsadsading .env file")
	}

	connection := db.InitDB()
	db.Transfer(connection)
	socketServer := socket.SocketInit()
	defer socketServer.Close()
	app := server.NewServer(connection)
	server.ConfigureRoutes(app)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
