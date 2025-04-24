package notifier

import (
	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/services"
)

type emailNotifier struct {
	EmailService domain.EmailServicePort
}

func NewEmailNotifier() domain.NotifierPort {
	return &emailNotifier{
		EmailService: services.NewEmailService(),
	}
}

func (n *emailNotifier) SendMessage(to []string, body string) {
	n.EmailService.SendMail(to, []byte(body))
}
