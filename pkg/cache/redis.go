package cache

import (
	"context"
	"goclean/config"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(cfg *config.Configuration) Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})

	redisClient.Ping(context.Background())

	return &RedisCache{redisClient}
}
