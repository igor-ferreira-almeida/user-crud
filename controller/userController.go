package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/domain/usermd"
	"github.com/igor-ferreira-almeida/user-crud/dto"
	"github.com/igor-ferreira-almeida/user-crud/service/usersvc"
	"strconv"

	"net/http"
)

var userService = usersvc.Inject()

func MapUserRoutes(router *gin.Engine) {
	router.GET("/users", FindUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
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

	addedUser := userService.Add(usermd.NewUser(userDTO.Name, userDTO.Age, userDTO.Gender))
	responseDTO := dto.NewUserResponseDTO(addedUser)

	context.JSON(http.StatusCreated, responseDTO)
}

// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new user
// @Success 200 {object} dto.UserDTO
// @Router /users [post]
func UpdateUser(context *gin.Context) {
	param := context.Param("id")
	userDTO := dto.UserDTO{}
	context.BindJSON(&userDTO)

	user := userDTO.ToModel()
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "Invalid param for id")
		context.JSON(http.StatusBadRequest, errorResponseDTO)
		return
	}

	updatedUser := userService.Update(id, user)

	userResponseDTO := dto.NewUserResponseDTO(updatedUser)
	context.JSON(http.StatusOK, userResponseDTO)
}

