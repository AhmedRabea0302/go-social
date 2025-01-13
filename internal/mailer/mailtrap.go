package mailer

import (
	"bytes"
	"errors"

	"text/template"

	"github.com/AhmedRabea0302/go-social/internal/env"
	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// Template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())

	message.AddAlternative("text/html", body.String())

	mailtrapUsername := env.GetString("MAILTRAP_USERNAME_KEY", "")
	mailtrapPassword := env.GetString("MAILTRAP_PASSWORD_KEY", "")

	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, mailtrapUsername, mailtrapPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return 200, nil
	}

	return 200, nil
}
