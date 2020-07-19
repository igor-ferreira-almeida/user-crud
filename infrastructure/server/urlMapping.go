package server

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/controller"
)

func mapRoutes(router *gin.Engine) {
	controller.MapHelloWorldRoutes(router)
	controller.MapUserRoutes(router)
}
