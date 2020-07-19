package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/dto"

	"net/http"
)

func MapUserRoutes(router *gin.Engine) {
	router.GET("/users", FindUser)
	router.POST("/users", CreateUser)
}

func FindUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Person finded!",
	})
}

func CreateUser(context *gin.Context) {
	userDTO := dto.UserDTO{}
	context.BindJSON(&userDTO)

	bytes, _ := json.Marshal(userDTO)
	fmt.Println(string(bytes))
	fmt.Println("-------------------------------")
	fmt.Println(userDTO)


	context.JSON(http.StatusCreated, userDTO)
}