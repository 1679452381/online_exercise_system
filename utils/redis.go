package utils

import (
	"github.com/redis/go-redis/v9"
)

var Redis = RedisClient()

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
