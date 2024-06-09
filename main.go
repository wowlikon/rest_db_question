package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

const (
	OK         = http.StatusOK
	NotFound   = http.StatusNotFound
	BadRequest = http.StatusBadRequest
)

func NewID() string {
	id := rand.Int63()
	return strconv.FormatInt(id, 10)
}

var addresses map[string]Address
var logger *log.Logger

func main() {
	addresses = make(map[string]Address)

	// Create a log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Create a logger that logs to both the file and the console
	logger = log.New(logFile, "", log.LstdFlags)
	logger.Println("Starting server...")
	r := gin.Default()
	r.Use(errorHandler)

	r.POST("/address", func(c *gin.Context) {
		id, err := CreateAddress(c)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(OK, gin.H{"id": id})
	})

	r.GET("/address/:id", func(c *gin.Context) {
		id := c.Param("id")
		addr, ok := GetAddressByID(id)
		if !ok {
			c.JSON(NotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(OK, addr)
	})

	logger.Println("Server started successfully!")
	r.Run(":80")
}

func CreateAddress(c *gin.Context) (string, error) {
	var addr Address
	var id string

	if err := c.Bind(&addr); err != nil {
		return "", err
	}

	for {
		id := NewID()
		if _, ok := addresses[id]; !ok {
			break
		}
	}
	addresses[id] = addr
	return id, nil
}

func GetAddressByID(id string) (Address, bool) {
	addr, ok := addresses[id]
	return addr, ok
}

func errorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		logger.Println("Error occurred:", c.Errors[0].Error())
		c.JSON(BadRequest, gin.H{"error": c.Errors[0].Error()})
	}
}
