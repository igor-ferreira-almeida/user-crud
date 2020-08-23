package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUser(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	FindUser(context)

	assert.EqualValues(t, response.Code, http.StatusOK)
}

func TestCreateUser(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	var jsonStr = []byte(`{"name":"Abel", "age":13, "gender":"male"}`)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	context.Request = request

	CreateUser(context)

	var responseDTO dto.UserResponseDTO
	err := json.Unmarshal(response.Body.Bytes(), &responseDTO)

	assert.Nil(t, err)
	assert.NotNil(t, responseDTO)
	assert.EqualValues(t, response.Code, http.StatusCreated)
	assert.EqualValues(t, "Abel", responseDTO.Name)
	assert.EqualValues(t, 13, responseDTO.Age)
	assert.EqualValues(t, "male", responseDTO.Gender)
}

func TestUpdateUser(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Params = gin.Params{
		{Key: "id", Value: "1"},
	}
	var jsonStr = []byte(`{"name":"Melissa", "age":16, "gender":"female"}`)
	request, _ := http.NewRequest("PUT", "/users/:id", bytes.NewBuffer(jsonStr))
	context.Request = request

	UpdateUser(context)

	var responseDTO dto.UserResponseDTO
	err := json.Unmarshal(response.Body.Bytes(), &responseDTO)

	assert.Nil(t, err)
	assert.NotNil(t, responseDTO)
	assert.EqualValues(t, response.Code, http.StatusOK)
	assert.EqualValues(t, 1, responseDTO.ID)
	assert.EqualValues(t, "Melissa", responseDTO.Name)
	assert.EqualValues(t, 16, responseDTO.Age)
	assert.EqualValues(t, "female", responseDTO.Gender)
}

func TestUpdateUserBadIdParam(t *testing.T) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	context.Params = gin.Params{
		{Key: "id", Value: "adfadsf"},
	}
	var jsonStr = []byte(`{"name":"Melissa", "age":16, "gender":"female"}`)
	request, _ := http.NewRequest("PUT", "/users/:id", bytes.NewBuffer(jsonStr))
	context.Request = request

	UpdateUser(context)

	var responseDTO dto.ErrorResponseDTO
	err := json.Unmarshal(response.Body.Bytes(), &responseDTO)

	assert.Nil(t, err)
	assert.NotNil(t, responseDTO)
	assert.EqualValues(t, response.Code, http.StatusBadRequest)
	assert.EqualValues(t, responseDTO.HttpStatus, http.StatusText(http.StatusBadRequest))
	assert.EqualValues(t, responseDTO.Code, http.StatusBadRequest)
	assert.EqualValues(t, responseDTO.Message, "Invalid param for id")
}
