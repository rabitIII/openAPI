package global

import (
	"rabit-api-backend/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
	Rdb    *redis.Client
)
