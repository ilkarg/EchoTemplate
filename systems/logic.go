package systems

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/go-playground/validator"
	"gopkg.in/gomail.v2"
)

func PasswordHashing(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func SendEmail(to, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "ilya.kargapolov02@mail.ru")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer("smtp.mail.ru", 587, "ilya.kargapolov02@mail.ru", "mQJvZ1WbGrMnGC4EcnWd")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func ValidateData(data interface{}) bool {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return false
	}
	return true
}
