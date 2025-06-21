package main

import (
	"go_web_test1/config"
	"go_web_test1/global"
	"go_web_test1/initialize"
	"go_web_test1/routes"
)

func main() {
	global.Config = config.InitConfig()

	// fmt.Printf("%+v\n", config.Appconfig)
	log, err := initialize.InitLogger("production", "./log/app.log")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	global.DB = initialize.InitDB()
	// global.Redis = initialize.InitRedis()

	r := routes.InitRouter()
	port := global.Config.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port)
}
