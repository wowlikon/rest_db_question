package api

import (
	"log"
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

type AddressAPI struct {
	Addresses map[string]Address
	Logger    *log.Logger
}

const (
	OK         = http.StatusOK
	NotFound   = http.StatusNotFound
	BadRequest = http.StatusBadRequest
)

// Добавление методов работы с базой данных
func InitMethods(r *gin.Engine, a AddressAPI) {
	r.POST("/address", func(c *gin.Context) {
		id, err := CreateAddress(c, a)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(OK, gin.H{"id": id})
	})

	r.GET("/address/:id", func(c *gin.Context) {
		id := c.Param("id")
		addr, ok := GetAddressByID(id, a)
		if !ok {
			c.JSON(NotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(OK, addr)
	})
}

// Генерация идентификатора
func NewID() string {
	id := rand.Int63()
	return strconv.FormatInt(id, 10)
}

// Операция добавления записи в базу данных
func CreateAddress(c *gin.Context, a AddressAPI) (string, error) {
	var addr Address
	var id string

	if err := c.Bind(&addr); err != nil {
		return "", err
	}

	for {
		id = NewID()
		if _, ok := a.Addresses[id]; !ok {
			break
		}
	}
	a.Addresses[id] = addr
	return id, nil
}

// Операция получения записи из базы данных
func GetAddressByID(id string, a AddressAPI) (Address, bool) {
	addr, ok := a.Addresses[id]
	return addr, ok
}

// Middleware для логирования ошибок
func NewErrorHandler(a AddressAPI) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			a.Logger.Println("Error occurred:", c.Errors[0].Error())
			c.JSON(BadRequest, gin.H{"error": c.Errors[0].Error()})
		}
	}
}
