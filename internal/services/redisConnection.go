package services

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitConnection() (context.Context, *redis.Client) {
	var ctx = context.Background()
	var redisClient *redis.Client = redis.NewClient(&redis.Options{
		Addr: utils.RedisHost,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(utils.ErrorFailedRedisConn, err)
	} else {
		fmt.Println(utils.SuccessRedisConn)
	}
	return ctx, redisClient
}
