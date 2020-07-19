package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapHelloWorldRoutes(router *gin.Engine) {
	router.POST("/hello-world", Welcome)
}

func Welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
