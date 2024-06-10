package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wneessen/go-mail"
	"github.com/wowlikon/rest_db_question/api"
)

var a api.AddressAPI

func main() {
	// Загружаем .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Создание файла для логов
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Подготовка AddressAPI. Создание логгерп
	logger := log.New(logFile, "", log.LstdFlags)
	a.Addresses = make(map[string]api.Address)
	a.Logger = logger

	// Создание mail клиента
	a.Mail, err := mail.NewClient("smtp.mail.ru",
		mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("bot_email")), mail.WithPassword(os.Getenv("bot_password")))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}

	// Инициализвция gin сервера
	logger.Println("Starting server...")
	r := gin.Default()

	r.Use(api.NewErrorHandler(a))
	api.InitMethods(r, a)

	logger.Println("Server started successfully!")
	r.Run(":80") // Запуск сервера на http://localhost:80

	defer logFile.Close()
}
