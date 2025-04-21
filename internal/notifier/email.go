package notifier

import (
	"github.com/JscorpTech/notification/internal/domain"
	"github.com/k0kubun/pp/v3"
)

type emailNotifier struct{}

func NewEmailNotifier() domain.NotifierPort {
	return &emailNotifier{}
}

func (n *emailNotifier) SendMessage(to []string, body string) {
	pp.Print(to, body)
}
