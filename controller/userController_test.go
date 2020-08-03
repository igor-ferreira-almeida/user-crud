package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUser(t *testing.T) {
	tests := []struct {
		tag            string
		expectedStatus int
	}{
		{
			tag:            "t1 find exists user",
			expectedStatus: http.StatusOK,
		},
	}
	for _, test := range tests {
		r := gin.Default()
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		rr := httptest.NewRecorder()
		r.GET("/users", FindUser)
		r.ServeHTTP(rr, req)

		var userDTO dto.UserDTO
		err := json.Unmarshal(rr.Body.Bytes(), &userDTO)


		assert.Nil(t, err)
		assert.NotNil(t, userDTO)
		assert.EqualValues(t, test.expectedStatus, rr.Code)
	}

}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		tag            string
		expectedStatus int
	}{
		{
			tag:            "t1 create valid user",
			expectedStatus: http.StatusCreated,
		},
	}

	for _, test := range tests {
		var jsonStr = []byte(`{"name":"Igor", "age":31, "gender":"male"}`)
		router := gin.Default()
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		router.POST("/users", CreateUser)
		router.ServeHTTP(rr, req)

		var userDTO dto.UserDTO
		err = json.Unmarshal(rr.Body.Bytes(), &userDTO)

		fmt.Println(userDTO)

		assert.Nil(t, err)
		assert.NotNil(t, userDTO)
		assert.EqualValues(t, test.expectedStatus, rr.Code)
		assert.EqualValues(t, "Igor", userDTO.Name)
		assert.EqualValues(t, 31, userDTO.Age)
		assert.EqualValues(t, "male", userDTO.Gender)
	}
}
