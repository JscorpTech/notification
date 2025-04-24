package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/JscorpTech/notification/internal/domain"
)

type emailService struct{}

func NewEmailService() domain.EmailServicePort {
	return &emailService{}
}

func (e *emailService) SendMail(to []string, body []byte) {
	// Gmail konfiguratsiyasi
	from := os.Getenv("MAIL_USER")
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_DOMAIN")
	smtpPort := os.Getenv("MAIL_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body)
	if err != nil {
		fmt.Println("Xatolik:", err)
		return
	}

	fmt.Println("Email yuborildi!")
}
