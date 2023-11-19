package service

import (
	"context"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	APIKey string
	Domain string
}

func (s *MailgunService) SendMail(from string, to []string, subject, body string) error {
	mg := mailgun.NewMailgun(s.Domain, s.APIKey)
	message := mg.NewMessage(
		from,
		subject,
		"",
		to...,
	)
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	return err
}
