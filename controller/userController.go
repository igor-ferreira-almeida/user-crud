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
	router.GET("/users/:id", FindUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
}

func FindUser(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "Invalid param for id")
		context.JSON(http.StatusBadRequest, errorResponseDTO)
		return
	}

	foundUser, err := userService.Find(id)

	if err != nil {
		errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusNotFound), http.StatusNotFound, "User not found")
		context.JSON(http.StatusNotFound, errorResponseDTO)
		return
	}

	userResponseDTO := dto.NewUserResponseDTO(foundUser)
	context.JSON(http.StatusOK, userResponseDTO)
}

// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new user
// @Success 200 {object} dto.UserRequestDTO
// @Router /users [post]
func CreateUser(context *gin.Context) {
	userDTO := dto.UserRequestDTO{}
	context.BindJSON(&userDTO)

	var x = context.Request.Header["X-Discount-Token"]
	fmt.Println(x)

	bytes, _ := json.Marshal(userDTO)
	fmt.Println(string(bytes))

	var addedUser, _ = userService.Add(usermd.NewUser(userDTO.Name, userDTO.Age, userDTO.Gender))
	responseDTO := dto.NewUserResponseDTO(addedUser)

	context.JSON(http.StatusCreated, responseDTO)
}

// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Summary Create a new user
// @Success 200 {object} dto.UserRequestDTO
// @Router /users [post]
func UpdateUser(context *gin.Context) {
	param := context.Param("id")
	userDTO := dto.UserRequestDTO{}
	context.BindJSON(&userDTO)

	user := userDTO.ToModel()
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusBadRequest), http.StatusBadRequest, "Invalid param for id")
		context.JSON(http.StatusBadRequest, errorResponseDTO)
		return
	}

	updatedUser, err := userService.Update(id, user)

	if err != nil {
		errorResponseDTO := dto.NewErrorResponseDTO(http.StatusText(http.StatusNotFound), http.StatusNotFound, "User not found")
		context.JSON(http.StatusNotFound, errorResponseDTO)
		return
	}

	userResponseDTO := dto.NewUserResponseDTO(updatedUser)
	context.JSON(http.StatusOK, userResponseDTO)
}

