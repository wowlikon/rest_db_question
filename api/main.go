package api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wneessen/go-mail"
)

type Address struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type AddressAPI struct {
	Addresses map[string]Address
	Mail      *mail.Client
	Logger    *log.Logger
	AdminMail string
	BotMail   string
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

	r.GET("/address/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		addr, ok := GetAddressByID(id, a)
		if !ok {
			c.JSON(NotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(OK, addr)
	})

	r.GET("/address/name/:name", func(c *gin.Context) {
		name := c.Param("name")
		addr, ok := GetAddressByName(name, a)
		if !ok {
			c.JSON(NotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(OK, gin.H{"result": addr, "count": len(addr)})
	})
}

// Генерация идентификатора
func NewID(obj map[string]Address) string {
	var id int64
	result := ""

	for {
		id = rand.Int63()
		result = strconv.FormatInt(id, 10)
		if _, ok := obj[result]; !ok {
			break
		}
	}
	return result
}

// Операция добавления записи в базу данных
func CreateAddress(c *gin.Context, a AddressAPI) (string, error) {
	var addr Address
	var id string

	if err := c.Bind(&addr); err != nil {
		return "", err
	}

	id = NewID(a.Addresses)
	if strings.HasPrefix(addr.Name, "Москва") {
		msgText := fmt.Sprintf(EmailInfo, addr.Name, addr)
		a.Logger.Print(msgText)
		send(a, msgText)

		addr.Name = "БЮ711"
	}

	a.Addresses[id] = addr
	return id, nil
}

// Операция получения записи из базы данных по id
func GetAddressByID(id string, a AddressAPI) (Address, bool) {
	addr, ok := a.Addresses[id]
	return addr, ok
}

// Операция получения записи из базы данных по name
func GetAddressByName(name string, a AddressAPI) ([]Address, bool) {
	res := []Address{}
	for _, element := range a.Addresses {
		if element.Name == name {
			res = append(res, element)
		}
	}

	return res, len(res) > 0
}

// Middleware для логирования ошибок
func NewErrorHandler(a AddressAPI) func(c *gin.Context) {
	return func(c *gin.Context) {
		a.Logger.Print(c.ClientIP(), c.Request.Method, c.Request.URL, "")

		start := time.Now()
		c.Next()
		a.Logger.Println(time.Since(start))

		if len(c.Errors) > 0 {
			a.Logger.Println("Error occurred:", c.Errors[0].Error())
			c.JSON(BadRequest, gin.H{"error": c.Errors[0].Error()})
		}
	}
}
