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
	userDTO := dto.UserDTO{Name: "name1", Age: 1, Gender: "female"}
	context.JSON(http.StatusOK, userDTO)
}

// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new user
// @Success 200 {object} dto.UserDTO
// @Router /users [post]
func CreateUser(context *gin.Context) {
	userDTO := dto.UserDTO{}
	context.BindJSON(&userDTO)

	var x = context.Request.Header["X-Discount-Token"]
	fmt.Println(x)

	bytes, _ := json.Marshal(userDTO)
	fmt.Println(string(bytes))
	fmt.Println("-------------------------------")
	fmt.Println(userDTO)


	context.JSON(http.StatusCreated, userDTO)
}