package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewID() string {
	id := rand.Int63()
	return strconv.FormatInt(id, 10)
}

var addresses map[string]Address

func main() {
	addresses = make(map[string]Address)
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

	r.GET("/address/:id", func(c *gin.Context) {
		addr, ok := addresses[c.Param("id")]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		c.JSON(http.StatusOK, addr)
	})

	r.Run()
}
