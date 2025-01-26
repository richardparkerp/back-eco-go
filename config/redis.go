package config

import (
	"context"
	"fmt"
	"os"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func InitRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})
}


