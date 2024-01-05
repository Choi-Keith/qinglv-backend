package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type SMTPOptions struct {
	Host string
	Port int
	User string
	Pass string
}

func Send(smtp SMTPOptions, toUser, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtp.User)
	m.SetHeader("To", toUser)
	m.SetHeader("Subject", subject)
	fmt.Printf("Send, smtp:%+v, toUser:%+v, subject: %+v\n", smtp, toUser, content)

	m.SetBody("text/html", content)
	d := gomail.NewDialer(smtp.Host, smtp.Port, smtp.User, smtp.Pass)
	return d.DialAndSend(m)
}
