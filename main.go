package main

import (
	"log"

	"github.com/setkyar/go-mail/config"
)

func main() {
	cfg := config.LoadMailConfig()
	mailService, err := config.NewMailService(cfg)

	if err != nil {
		log.Fatalf("Failed to initialize mail service: %s", err)
	}

	err = mailService.SendMail("from@setkyar.com", []string{"me@setkyar.com"}, "Testing: Go Mail", "Hey, this is a test mail!")
	if err != nil {
		log.Fatalf("Failed to send mail: %s", err)
	}

	log.Println("Mail sent successfully")
}
