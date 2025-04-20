package notifier

import (
	"github.com/JscorpTech/notification/internal/domain"
	"github.com/k0kubun/pp/v3"
)

type smsNotifier struct{}

func NewSmsNotifier() domain.NotifierPort {
	return &smsNotifier{}
}

func (n *smsNotifier) SendMessage(to []string, body string) {
	pp.Print(to, body)
}
