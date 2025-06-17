package initialize

import (
	"go_web_test1/global"
	"os"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.Db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		global.Log.Error("Failed to connect to Redis:", zap.Error(err))
		os.Exit(1)
	}

	return *client
}
