package domain

import "github.com/streadway/amqp"

type NotificationConsumerPort interface {
	Start()
	Handler(amqp.Delivery)
}

type SMSServicePort interface {
	SendSMS(string, string) error
}
type NotifierPort interface {
	SendMessage([]string, string)
}

type NotificationMsg struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	To      []string `json:"to"`
}
