package api

import (
	"github.com/wneessen/go-mail"
)

const (
	EmailSubject = "Замена name в API сервера"
	EmailInfo    = "Replaced \"%s\" to \"БЮ711\" in %v\n"
)

func send(a AddressAPI, text string) {
	// Создание сообщения
	m := mail.NewMsg()
	if err := m.From(a.BotMail); err != nil {
		a.Logger.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(a.AdminMail); err != nil {
		a.Logger.Fatalf("failed to set To address: %s", err)
	}

	m.Subject(EmailSubject)
	m.SetBodyString(
		mail.TypeTextPlain,
		text,
	)

	// Отправка сообщения
	if err := a.Mail.DialAndSend(m); err != nil {
		a.Logger.Fatalf("failed to send mail: %s", err)
	}
}
