package main

import (
	"os"
	"github.com/joho/godotenv"

	"github.com/YanaPIIDXer/QiitaLineBot/src/infrastructure/connection"
)

func main() {
	godotenv.Load(".env")

	var connection = InfrastructureConnection.NewHTTPConnection()
	connection.AddAccessEvent(InfrastructureConnection.NewLocalAccessEvent())
	connection.AddAccessEvent(InfrastructureConnection.NewLINEAccessEvent())
	
	var port = os.Getenv("PORT")		// heroku上で動かす場合は指定されたPortじゃないと死なますよ。
	if port == "" {
		port = "3000"
	}
	connection.Service(port)
}
