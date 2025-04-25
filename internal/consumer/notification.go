package consumer

import (
	"context"
	"fmt"
	"os"

	"github.com/JscorpTech/notification/internal/broker"
	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/notifier"
)

type notificationConsumer struct {
	Ctx context.Context
}

func NewNotificationConsumer(ctx context.Context) domain.NotificationConsumerPort {
	return &notificationConsumer{
		Ctx: ctx,
	}
}

func (n *notificationConsumer) Start() {

	brokerName := os.Getenv("BROKER")
	if brokerName == "" {
		brokerName = "redis"
	}
	var brokerService domain.BrokerPort
	switch brokerName {
	case "redis":
		brokerService = broker.NewRedisBroker(n.Ctx)
	case "rabbitmq":
		brokerService = broker.NewRabbitMQBroker(n.Ctx)
	default:
		brokerService = broker.NewRedisBroker(n.Ctx)
	}
	brokerService.Subscribe(os.Getenv("TOPIC"), n.Handler)
	fmt.Println("🚀 Server started. Ctrl+C to quit.")
	select {}
}

func (n *notificationConsumer) Handler(notification domain.NotificationMsg) {
	var ntf domain.NotifierPort
	switch notification.Type {
	case "sms":
		ntf = notifier.NewSmsNotifier(n.Ctx)
	case "email":
		ntf = notifier.NewEmailNotifier()
	}
	ntf.SendMessage(notification.To, notification.Message)
}
