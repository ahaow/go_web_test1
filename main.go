package main

import (
	"go_web_test1/config"
	"go_web_test1/global"
	"go_web_test1/logger"
	"go_web_test1/routes"
)

func main() {
	// 1. 初始化配置
	global.Config = config.InitConfig()
	// fmt.Printf("%+v\n", config.Appconfig)

	if global.Config == nil {
		panic("配置加载失败")
	}

	// 2. 初始化日志
	log, err := logger.NewLogger("development", "logs", "[myApp] ")
	if err != nil {
		panic(err)
	}
	global.Log = log

	defer log.Close() // 确保程序结束时同步日志

	global.Log.Info("配置与日志初始化完成")

	// 3. 初试化数据库
	// global.DB = initialize.InitDB()
	// if global.DB == nil {
	// 	global.Log.Fatal("数据库初始化失败")
	// }
	// global.Redis = initialize.InitRedis()

	r := routes.InitRouter()
	port := global.Config.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port)

}
