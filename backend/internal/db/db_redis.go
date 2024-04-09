package db

import (
	"backend-go/internal/conf"
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.Addr(),
		Password: conf.Conf.Redis.Password,
		DB:       0,
	})

	Rdb = rdb
	return
}
