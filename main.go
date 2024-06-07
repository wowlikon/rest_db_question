
package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AddressData struct {
	name string `json:"name"`
	address string `json: "address"`
	longitude float64 `json: "longitude"`
	latitude float64 `json: "latitude"`
}

func NewID() json.Number {
	id := rand.Int63()
	return json.Number(strconv.FormatInt(id, 10))
}

func main() {
	addresses := make(map[json.Number]Address)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(NewID())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
    		c.JSON(http.StatusOK, gin.H{
    		  "message": "pong",
    		})
  	})
	r.POST("/address", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"id": "number",
		})
	})
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
