package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func GenerateResetToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendResetEmail(smtpConfig SMTPConfig, email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", smtpConfig.From)
	m.SetHeader("To", email)
	m.SetBody("text/plain", fmt.Sprintf("Here is your password reset link: http://localhost:3000/newpass?token=%s", token))

	d := gomail.NewDialer(smtpConfig.Host, smtpConfig.Port, smtpConfig.Username, smtpConfig.Password)
	return d.DialAndSend(m)
}
