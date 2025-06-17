package global

import (
	"go_web_test1/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *zap.Logger
	Redis  redis.Client
)
