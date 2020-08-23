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

//func TestCreateUser(t *testing.T) {
//	tests := []struct {
//		tag            string
//		expectedStatus int
//	}{
//		{
//			tag:            "t1 create valid user",
//			expectedStatus: http.StatusCreated,
//		},
//	}
//
//	for _, test := range tests {
//		var jsonStr = []byte(`{"name":"Igor", "age":31, "gender":"male"}`)
//		router := gin.Default()
//		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
//
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		req.Header.Set("Content-Type", "application/json")
//		rr := httptest.NewRecorder()
//
//		// Mock
//		//user := usermd.User{"name1", 1, "male"}
//		//mockService := new(MockService)
//		//mockService.On("Add").Return(user)
//		// userService := mockService
//		// userService.Add(user)
//
//		router.POST("/users", CreateUser)
//		router.ServeHTTP(rr, req)
//
//		var userDTO dto.UserDTO
//		err = json.Unmarshal(rr.Body.Bytes(), &userDTO)
//
//		fmt.Println(userDTO)
//
//		assert.Nil(t, err)
//		assert.NotNil(t, userDTO)
//		assert.EqualValues(t, test.expectedStatus, rr.Code)
//		assert.EqualValues(t, "Igor", userDTO.Name)
//		assert.EqualValues(t, 31, userDTO.Age)
//		assert.EqualValues(t, "male", userDTO.Gender)
//	}
//}

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
