package routes

import (
	"go_web_test1/controllers"
	"go_web_test1/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.GET("/register", controllers.Register)
		auth.GET("/login", controllers.Login)
	}

	user := r.Group("/api/user")
	user.Use(middleware.AuthJWT())
	{
		user.GET("profile", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}
}
