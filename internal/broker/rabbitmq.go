package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/rabbitmq"
)

type rabbitMQBroker struct {
	Ctx context.Context
}

func NewRabbitMQBroker(ctx context.Context) domain.BrokerPort {
	return &rabbitMQBroker{
		Ctx: ctx,
	}
}

func (r rabbitMQBroker) Subscribe(topic string, handler func(domain.NotificationMsg)) {
	conn, ch, err := rabbitmq.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	ch.ExchangeDeclare(topic, "direct", true, false, false, false, nil)
	q, _ := ch.QueueDeclare(topic, true, false, false, false, nil)
	ch.QueueBind(q.Name, topic, topic, false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for msg := range msgs {
			var notification domain.NotificationMsg
			if err := json.Unmarshal(msg.Body, &notification); err != nil {
				fmt.Print(err.Error())
			}
			go handler(notification)
		}
	}()
}
