package server

import (
	"github.com/gin-gonic/gin"
)

// Start is a function to start server
func Start() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	mapRoutes(router)
	router.Run(":9080")
}
