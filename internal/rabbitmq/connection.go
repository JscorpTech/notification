package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
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
