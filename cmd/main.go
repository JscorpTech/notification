package main

import (
	"context"

	"github.com/JscorpTech/notification/internal/consumer"
	"github.com/JscorpTech/notification/internal/redis"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	redis.InitRedis()
	notification := consumer.NewNotificationConsumer(ctx)
	notification.Start()
}
