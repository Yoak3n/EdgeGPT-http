package database

import (
	"github.com/go-redis/redis/v8"
	"time"
)

func NewRedis() {
	redis.NewClient(&redis.Options{
		PoolSize:    10,
		PoolTimeout: time.Second * 5,
	})
}
