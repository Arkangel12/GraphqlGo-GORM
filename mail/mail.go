package mail

import (
	"github.com/Arkangel12/curso/graphql-gorm/config"
	"log"
	"net/smtp"
)

type Mail struct {
	Name    string
	Email   string
	Subject string
	Body    string
}

func Send(email Mail) error {
	from := config.AppConfig.Email
	pass := config.AppConfig.EPassword
	to := email.Email

	msg := "From: " + config.AppConfig.Name + "\n" +
		"To: " + email.Email + "\n" +
		"Subject: " + email.Subject + "\n\n" +
		email.Body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print("Your message has been sent!")
	return nil
}
