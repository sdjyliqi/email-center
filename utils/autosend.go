package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

var emailAuth = smtp.PlainAuth("", "sdjyliqi@163.com", "liqidxq753129", "smtp.163.com")

//SendToMail ...
func SendToMail(user, to, subject string, text []byte) error {
	e := &email.Email{
		To:      []string{to},
		From:    user,
		Subject: subject,
		Text:    text,
		HTML:    []byte(""),
		Headers: textproto.MIMEHeader{},
	}
	return e.Send(SMTP163Host, emailAuth)
}
