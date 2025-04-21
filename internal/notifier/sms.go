package notifier

import (
	"context"

	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/services"
)

type smsNotifier struct {
	SMSServie domain.SMSServicePort
	Ctx       context.Context
}

func NewSmsNotifier(ctx context.Context) domain.NotifierPort {
	return &smsNotifier{
		SMSServie: services.NewEskizSMSService(ctx),
	}
}

func (n *smsNotifier) SendMessage(to []string, body string) {
	for _, user := range to {
		n.SMSServie.SendSMS(user, body)
	}
}
