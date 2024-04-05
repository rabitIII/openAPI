package db

import (
	"backend-go/config"
	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Addr(),
		Password: config.Conf.Redis.Password,
		DB:       0,
	})

	Rdb = rdb
	return
}
