package studentcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirio-neto/gin-rest-api/database"
	studentmodel "github.com/sirio-neto/gin-rest-api/models/StudentModel"
)

func Greet(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"message": "Bem vindo " + nome + "!",
	})
}

func GetByCpf(c *gin.Context) {
	var student studentmodel.Student

	cpf := c.Params.ByName("cpf")
	database.DB.Where(&studentmodel.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Student not found.",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func GetAll(c *gin.Context) {
	var students []studentmodel.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	var student studentmodel.Student
	database.DB.First(&student, parsedId)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Student not found.",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func Insert(c *gin.Context) {
	var student studentmodel.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	var student studentmodel.Student
	database.DB.First(&student, parsedId)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Student not found.",
		})

		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := student.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Save(&student)

	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	var student studentmodel.Student
	database.DB.Delete(&student, parsedId)

	c.JSON(http.StatusOK, gin.H{})
}
