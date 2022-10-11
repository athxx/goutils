package utils

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

var (
	eml     *email.Email
	emlAddr string
	emlAuth smtp.Auth
)

func EmailInit(addr, identity, user, psw, host, from string) {
	eml = email.NewEmail()
	eml.From = from + " <" + user + ">"
	emlAddr = addr
	emlAuth = smtp.PlainAuth(identity, user, psw, host)
}

func EmailSend(subject, content string, to []string) error {
	eml.To = to
	eml.Subject = subject
	eml.HTML = []byte(content)
	return eml.Send(emlAddr, emlAuth)
}
