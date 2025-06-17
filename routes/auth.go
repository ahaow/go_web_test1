package routes

import (
	"go_web_test1/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.GET("/register", controllers.Register)
		auth.GET("/login", controllers.Login)
	}
}
