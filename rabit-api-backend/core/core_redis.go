package core

import (
	"fmt"
	"rabit-api-backend/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	fmt.Println("redis: ", global.Config.Databases.Redis.GetAddr())
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Databases.Redis.GetAddr(),
		Password: global.Config.Databases.Redis.Password, // no password set
		DB:       global.Config.Databases.Redis.DB,       // use default DB
	})

	fmt.Println("[Success] Redis数据库连接成功！！！")
	return client
}
