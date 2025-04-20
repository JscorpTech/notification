package main

import (
	"github.com/JscorpTech/notification/internal/consumer"
)

func main() {
	notification := consumer.NewNotificationConsumer()
	notification.Start()
}
