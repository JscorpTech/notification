package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/JscorpTech/notification/internal/domain"
	"github.com/JscorpTech/notification/internal/redis"
)

type redisBroker struct {
	Ctx context.Context
}

func NewRedisBroker(ctx context.Context) domain.BrokerPort {
	return &redisBroker{
		Ctx: ctx,
	}
}

func (r redisBroker) Subscribe(topic string, handler func(domain.NotificationMsg)) {
	go func() {
		for {
			var notification domain.NotificationMsg
			val, err := redis.RDB.BLPop(r.Ctx, 0, topic).Result()
			if err != nil {
				fmt.Print(err.Error())
				return
			}
			if err := json.Unmarshal([]byte(val[1]), &notification); err != nil {
				fmt.Print(err.Error())
				return
			}
			go handler(notification)
		}
	}()

}
