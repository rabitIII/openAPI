package core

import (
	"backend-go/global"
	"github.com/go-redis/redis"
)

func InitRedis(db int) *redis.Client {
	redisConfig := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr(),
		Password: redisConfig.Password,
		DB:       db,
		PoolSize: redisConfig.PoolSize,
	})

	//_, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
	//defer cancel()
	//_, err := client.Ping().Result()
	//if err != nil {
	//	logrus.Fatalf("%s redis连接失败，err:%s", redisConfig.Addr(), err.Error())
	//}
	return client
}
