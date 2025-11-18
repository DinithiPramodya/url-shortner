package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Memurai/Redis port
		DB:   0,                // Use default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("✅ Connected to Redis!")
}
