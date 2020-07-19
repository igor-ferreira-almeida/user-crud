package server

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func mapRoutes(router *gin.Engine) {
	url := ginSwagger.URL("http://localhost:9080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	controller.MapHelloWorldRoutes(router)
	controller.MapUserRoutes(router)
}
