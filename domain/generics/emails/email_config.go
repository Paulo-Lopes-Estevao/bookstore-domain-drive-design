package emails

import (
	"crypto/tls"
	"errors"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(to, Subject, body, contentType string) error {

	var from = os.Getenv("EMAIL_FROM")
	var password = os.Getenv("SMTP_PASS")
	var smtpHost = os.Getenv("SMTP_HOST")
	var smtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", Subject)
	m.SetBody(contentType, body)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return errors.New("Could not send email: " + err.Error())
	}

	return nil
}
