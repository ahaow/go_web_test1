package controllers

import (
	"go_web_test1/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(ctx *gin.Context) {

	global.Log.Error("Registration failed", zap.String("Username", "foo"))

	ctx.JSON(http.StatusOK, gin.H{"token": "123123"})

}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"token": "assss"})

}
