package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/setkyar/go-mail/service"
)

type MailConfig struct {
	ServiceType string
	SMTP        SMTPConfig
	Mailgun     MailgunConfig
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type MailgunConfig struct {
	APIKey string
	Domain string
	Secret string
}

func LoadMailConfig() *MailConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	return &MailConfig{
		ServiceType: os.Getenv("MAIL_SERVICE_TYPE"),
		SMTP: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
		},
		Mailgun: MailgunConfig{
			APIKey: os.Getenv("MAILGUN_API_KEY"),
			Domain: os.Getenv("MAILGUN_DOMAIN"),
		},
	}
}

func NewMailService(cfg *MailConfig) (service.MailService, error) {
	switch cfg.ServiceType {
	case "SMTP":
		return &service.SMTPService{
			Host:     cfg.SMTP.Host,
			Port:     cfg.SMTP.Port,
			Username: cfg.SMTP.Username,
			Password: cfg.SMTP.Password,
		}, nil
	case "Mailgun":
		return &service.MailgunService{
			APIKey: cfg.Mailgun.APIKey,
			Domain: cfg.Mailgun.Domain,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported mail service: %s", cfg.ServiceType)
	}
}
