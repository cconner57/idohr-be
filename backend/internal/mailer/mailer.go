package mailer

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	host     string
	port     int
	username string
	password string
	sender   string
}

func New(host string, port int, username, password, sender string) Mailer {
	return Mailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
		sender:   sender,
	}
}

func (m Mailer) Send(recipient, subject string, body string) error {
	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	addr := fmt.Sprintf("%s:%d", m.host, m.port)

	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n"+
		"\r\n"+
		"%s\r\n", m.sender, recipient, subject, body))

	return smtp.SendMail(addr, auth, m.sender, []string{recipient}, msg)
}
