package database

import (
	"log"

	studentmodel "github.com/sirio-neto/gin-rest-api/models/StudentModel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnDB() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Ocorreu um erro ao conectar com banco de dados", err)
	}

	DB.AutoMigrate(&studentmodel.Student{})
}
