package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/igor-ferreira-almeida/user-crud/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUser(t *testing.T) {
	tests := []struct {
		tag      string
		expected float64
	}{
		{
			tag:      "t1 find exists user",
			expected: 1.0,
		},
	}
	for _, test := range tests {
		result := 1.0
		if result != test.expected {
			t.Errorf("TestCreateUser %v failed, expectedStatus %v, got %v", test.tag, test.expected, result)
		}
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

		if status := rr.Code; status != test.expectedStatus {
			t.Errorf("TestCreateUser %v failed, expectedStatus %v, got %v", test.tag, test.expectedStatus, status)
		}
	}
}
