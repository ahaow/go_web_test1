package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 权限模块
	AuthRouter(r)

	return r
}
