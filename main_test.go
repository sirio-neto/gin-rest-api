package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	studentcontroller "github.com/sirio-neto/gin-rest-api/controllers/StudentController"
	"github.com/sirio-neto/gin-rest-api/database"
	studentmodel "github.com/sirio-neto/gin-rest-api/models/StudentModel"
	"github.com/stretchr/testify/assert"
)

var MockedStudentId int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return r
}

func TestGreet(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:nome", studentcontroller.Greet)

	tests := []string{"aluno", "aluna", "dev"}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "/"+test, nil)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, req)

		expected := "{\"message\":\"Bem vindo " + test + "!\"}"
		responseBody, _ := ioutil.ReadAll(response.Body)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expected, string(responseBody))
	}
}

func TestGreetFail(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:nome", studentcontroller.Greet)

	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func TestGetAll(t *testing.T) {
	database.ConnDB()
	mockStudent()

	r := SetupTestRoutes()
	r.GET("/students", studentcontroller.GetAll)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotNil(t, response.Body)

	deleteMockedStudent()
}

func TestGetByCpf(t *testing.T) {
	database.ConnDB()
	mockStudent()

	r := SetupTestRoutes()
	r.GET("/students/cpf/:cpf", studentcontroller.GetByCpf)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotNil(t, response.Body)

	deleteMockedStudent()
}

func mockStudent() {
	student := studentmodel.Student{
		Name: "Test Student",
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.DB.Create(&student)

	MockedStudentId = int(student.ID)
}

func deleteMockedStudent() {
	var student studentmodel.Student

	database.DB.Delete(&student, MockedStudentId)
}
