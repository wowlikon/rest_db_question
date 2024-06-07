
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Name string `json:"name"`
	Address string `json: "address"`
	Longitude float64 `json: "longitude"`
	Latitude float64 `json: "latitude"`
}

type Query struct {
	id string `json: "id"`
}

func NewID() string {
	id := rand.Int63()
	return strconv.FormatInt(id, 10)
}

var addresses map[string]Address

func main() {
	addresses = make(map[string]Address)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(NewID())

	r := gin.Default()
	r.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, addresses)
	})
	r.POST("/address", func(c *gin.Context) {
		var addr Address
		id := NewID()
		if err := c.Bind(&addr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		addresses[id] = addr
		c.JSON(http.StatusOK, gin.H{"id": id})
	})
	r.GET("/address", func(c *gin.Context) {
		var q Query
		if err := c.Bind(&q); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		addr, ok := addresses[q.id]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		
		c.JSON(http.StatusOK, addr)
	})
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
