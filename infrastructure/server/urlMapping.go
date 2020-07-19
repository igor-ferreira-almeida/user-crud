package server

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/igor-ferreira-almeida/user-crud/docs"
)

// @title User Crud API
// @version 1.0
// @description This is a generic user crud.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9080
// @BasePath /
func mapRoutes(router *gin.Engine) {
	url := ginSwagger.URL("http://localhost:9080/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	controller.MapHelloWorldRoutes(router)
	controller.MapUserRoutes(router)
}
