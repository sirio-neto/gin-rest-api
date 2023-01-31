package routes

import (
	"github.com/gin-gonic/gin"
	studentcontroller "github.com/sirio-neto/gin-rest-api/controllers/StudentController"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/:nome", studentcontroller.Greet)
	r.GET("/students", studentcontroller.GetAll)
	r.GET("/students/cpf/:cpf", studentcontroller.GetByCpf)
	r.GET("/students/:id", studentcontroller.GetById)
	r.DELETE("/students/:id", studentcontroller.Delete)
	r.PATCH("/students/:id", studentcontroller.Update)
	r.POST("/students", studentcontroller.Insert)

	r.Run(":8002")
}
