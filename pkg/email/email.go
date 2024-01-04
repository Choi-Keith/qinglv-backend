package email

import (
	"gopkg.in/gomail.v2"
)

func Send(toUser string, subject string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "1278028801@qq.com")
	m.SetHeader("To", toUser)
	m.SetHeader("Subject", subject)
	// fmt.Printf("Send, %+v, %+v\n", toUser, content)

	m.SetBody("text/html", content)
	d := gomail.NewDialer("xxx", 465, "xxx", "xxx")
	return d.DialAndSend(m)
}
