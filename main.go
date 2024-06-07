
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

type Address struct {
	Name string `json:"name"`
	Address string `json: "address"`
	Longitude float64 `json: "longitude"`
	Latitude float64 `json: "latitude"`
}

type Query struct {
	id json.Number `json: "id"`
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
	r.POST("/address", func(c *gin.Context) {
		var addr Address
		id := NewID()
		if err := c.Bind(&addr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		addresses[id] = addr
		c.JSON(http.StatusOK, Query{id})
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
