package middleware

import (
	"go_web_test1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthJWT 是 Gin 的中间件，用于校验每一个需要认证的API
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Authorization头中获取令牌
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "缺少认证令牌",
			})
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := utils.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "无效认证令牌",
				"error":   err.Error(), //可以打印具体错误以便调试
			})
			c.Abort()
			return
		}

		// 将Username保存到Context中，后续handler可以获取到
		c.Set("Username", claims.Username)

		// 继续处理
		c.Next()
	}
}
