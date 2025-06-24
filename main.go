package main

import (
	"fmt"
	"go_web_test1/logger"

	"go.uber.org/zap"
)

func main() {
	// global.Config = config.InitConfig()

	// // fmt.Printf("%+v\n", config.Appconfig)
	// log, err := initialize.InitLogger("production", "./logs/app.log")
	// if err != nil {
	// 	panic(err)
	// }
	// defer log.Close()

	// global.DB = initialize.InitDB()
	// // global.Redis = initialize.InitRedis()

	// r := routes.InitRouter()
	// port := global.Config.App.Port
	// if port == "" {
	// 	port = ":8080"
	// }
	// r.Run(port)

	log, err := logger.NewLogger("development", "logs", "[myApp] ")
	if err != nil {
		panic(err)
	}
	defer log.Close() // 确保程序结束时同步日志

	// 基本日志记录
	log.Debug("this is debug") // 仅开发环境可见
	log.Info("this is info")
	log.Warn("this is warning")
	log.Error("this is error")

	// 带字段的日志
	// 使用全局 logger
	zap.S().Infof("%s xxx", "carpe")

	// 示例：模拟多条错误日志
	for i := 1; i <= 3; i++ {
		log.Error(fmt.Sprintf("error %d", i))
	}
}
