package global

import (
	"rabit-api-backend/config"

	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
)
