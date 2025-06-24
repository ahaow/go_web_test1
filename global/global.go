package global

import (
	"go_web_test1/config"
	"go_web_test1/logger"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logger.Logger
	Redis  redis.Client
)
