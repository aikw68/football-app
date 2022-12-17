package util

import (
	"gopkg.in/gomail.v2"
)

// 作成中
func SendMail(email, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "from@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.Dialer{Host: "mailhog", Port: 1025}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
