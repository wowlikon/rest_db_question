package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wowlikon/rest_db_question/api"
)

var a api.AddressAPI

func main() {
	a.addresses = make(map[string]api.Address)

	// Create a log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Create a logger that logs to both the file and the console
	logger := log.New(logFile, "", log.LstdFlags)
	a.logger = logger

	logger.Println("Starting server...")
	r := gin.Default()

	r.Use(api.NewErrorHandler(a))
	api.InitMethods(r)

	logger.Println("Server started successfully!")
	r.Run(":80")

	defer logFile.Close()
}
