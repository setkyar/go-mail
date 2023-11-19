package service

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SMTPService struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (s *SMTPService) SendMail(from string, to []string, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	serverAddr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	message := buildMessage(s.Username, to, subject, body)

	err := smtp.SendMail(serverAddr, auth, s.Username, to, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func buildMessage(from string, to []string, subject, body string) string {
	header := fmt.Sprintf("From: %s\r\n", from)
	header += fmt.Sprintf("To: %s\r\n", strings.Join(to, ";"))
	header += fmt.Sprintf("Subject: %s\r\n", subject)
	header += "\r\n" // Blank line to separate headers from body

	return header + body
}
