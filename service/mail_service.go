package service

type MailService interface {
	SendMail(from string, to []string, subject, body string) error
}
