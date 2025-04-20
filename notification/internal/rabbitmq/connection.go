package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	log.Println("🐇 Connected to RabbitMQ")

	return conn, ch, nil
}
