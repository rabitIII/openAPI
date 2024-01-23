package global

import (
	"backend-go/config"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config // 全局指针
	Log    *logrus.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)
