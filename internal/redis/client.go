package redis

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRedis() {
	DB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       DB,
	})

	// Test connection
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redisga ulanib bo'lmadi: %v", err)
	}
}
