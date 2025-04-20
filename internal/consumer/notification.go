package consumer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/notifier"
	"github.com/JscorpTech/notification/internal/rabbitmq"
	"github.com/streadway/amqp"
)

type notificationConsumer struct{}

func NewNotificationConsumer() domain.NotificationConsumerPort {
	return &notificationConsumer{}
}

func (n *notificationConsumer) Start() {
	conn, ch, err := rabbitmq.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	exchangeName := "notification"
	queueName := "notification"
	routingKey := "notification"

	ch.ExchangeDeclare(exchangeName, "direct", true, false, false, false, nil)
	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	ch.QueueBind(q.Name, routingKey, exchangeName, false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for msg := range msgs {
			go n.Handler(msg)
		}
	}()

	fmt.Println("🚀 Server started. Ctrl+C to quit.")
	select {}
}

func (n *notificationConsumer) Handler(msg amqp.Delivery) {
	var notification domain.NotificationMsg
	err := json.Unmarshal(msg.Body, &notification)
	if err != nil {
		fmt.Print(err.Error())
	}
	var ntf domain.NotifierPort
	switch notification.Type {
	case "sms":
		ntf = notifier.NewSmsNotifier()
	case "email":
		ntf = notifier.NewEmailNotifier()
	}
	ntf.SendMessage(notification.To, notification.Message)
}
