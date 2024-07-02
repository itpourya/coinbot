package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type CacheRepository interface {
	PUSH(msg string)
	POP() (bool, string)
}

type cacheRepository struct {
	redis *redis.Client
}

var ctx = context.Background()

func (c cacheRepository) PUSH(msg string) {
	c.redis.RPush(ctx, "command", msg)
}

func (c cacheRepository) POP() (bool, string) {
	msg, err := c.redis.LPop(ctx, "command").Result()
	if err != nil {
		return false, err.Error()
	}

	return true, msg
}

func NewCacheRepository(connection *redis.Client) CacheRepository {
	return cacheRepository{redis: connection}
}
