package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	studentcontroller "github.com/sirio-neto/gin-rest-api/controllers/StudentController"
	"github.com/sirio-neto/gin-rest-api/database"
	studentmodel "github.com/sirio-neto/gin-rest-api/models/StudentModel"
	"github.com/stretchr/testify/assert"
)

var MockedStudent studentmodel.Student

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

func TestGetById(t *testing.T) {
	mockStudent()

	r := SetupTestRoutes()
	r.GET("/students/:id", studentcontroller.GetById)

	mockedId := strconv.Itoa(int(MockedStudent.ID))
	req, _ := http.NewRequest("GET", "/students/"+mockedId, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var student studentmodel.Student
	json.Unmarshal(response.Body.Bytes(), &student)

	assert.Equal(t, MockedStudent.ID, student.ID)
	assert.Equal(t, MockedStudent.Name, student.Name)
	assert.Equal(t, MockedStudent.CPF, student.CPF)
	assert.Equal(t, MockedStudent.RG, student.RG)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.NotNil(t, response.Body)

	deleteMockedStudent()
}

func TestDelete(t *testing.T) {
	mockStudent()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", studentcontroller.Delete)

	mockedId := strconv.Itoa(int(MockedStudent.ID))
	req, _ := http.NewRequest("DELETE", "/students/"+mockedId, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdate(t *testing.T) {
	mockStudent()

	r := SetupTestRoutes()
	r.PATCH("/students/:id", studentcontroller.Update)

	student := studentmodel.Student{
		Name: "Test Updated Student",
		CPF:  "12345678900",
		RG:   "123456700",
	}

	mockedId := strconv.Itoa(int(MockedStudent.ID))
	jsonParsed, _ := json.Marshal(student)

	req, _ := http.NewRequest("PATCH", "/students/"+mockedId, bytes.NewBuffer(jsonParsed))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var updatedStudent studentmodel.Student
	json.Unmarshal(response.Body.Bytes(), &updatedStudent)

	assert.Equal(t, student.Name, updatedStudent.Name)
	assert.Equal(t, student.CPF, updatedStudent.CPF)
	assert.Equal(t, student.RG, updatedStudent.RG)

	deleteMockedStudent()
}

func mockStudent() {
	MockedStudent = studentmodel.Student{
		Name: "Test Student",
		CPF:  "12345678901",
		RG:   "123456789",
	}

	database.ConnDB()
	database.DB.Create(&MockedStudent)
}

func deleteMockedStudent() {
	var student studentmodel.Student

	database.ConnDB()
	database.DB.Delete(&student, MockedStudent.ID)
}
