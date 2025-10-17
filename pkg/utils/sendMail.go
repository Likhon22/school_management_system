package utils

import (
	"sync"

	"github.com/go-mail/mail/v2"
)

var (
	once       sync.Once
	mailDialer *mail.Dialer
)

func GetMailDialer(host string, port int, username, password string) *mail.Dialer {
	once.Do(func() {
		mailDialer = mail.NewDialer(host, port, username, password)
	})
	return mailDialer

}

func SendMail(from string, to []string, subject, body string) error {

	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	return GetMailDialer("localhost", 1025, "", "").DialAndSend(m)
}
