package api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

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

	if strings.HasPrefix(addr.Name, "Москва") {
		addr.Name = "БЮ711"
		a.Logger.Printf("Replaced \"Москва\" to \"БЮ711\" in %v\n", addr)

		// Создание сообщения
		m := mail.NewMsg()
		if err := m.From(a.BotMail); err != nil {
			a.Logger.Fatalf("failed to set From address: %s", err)
		}
		if err := m.To(a.AdminMail); err != nil {
			a.Logger.Fatalf("failed to set To address: %s", err)
		}
		m.Subject("Замена name в API сервера")
		m.SetBodyString(
			mail.TypeTextPlain,
			fmt.Sprintf("Replaced \"Москва\" to \"БЮ711\" in %v\n", addr),
		)

		// Отправка сообщения
		if err := a.Mail.DialAndSend(m); err != nil {
			a.Logger.Fatalf("failed to send mail: %s", err)
		}
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
func GetAddressByName(name string, a AddressAPI) (Address, bool) {
	for _, element := range a.Addresses {
		if element.Name == name {
			return element, true
		}
	}
	return Address{}, false
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
