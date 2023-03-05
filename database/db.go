package database

import (
	"log"

	config "github.com/sirio-neto/gin-rest-api/environment"
	studentmodel "github.com/sirio-neto/gin-rest-api/models/StudentModel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnDB() {
	DB, err = gorm.Open(postgres.Open(getConnectionString()))

	if err != nil {
		log.Panic("Ocorreu um erro ao conectar com banco de dados", err)
	}

	DB.AutoMigrate(&studentmodel.Student{})
}

func getConnectionString() string {
	connString := "host=" + config.Env.DbHost
	connString += " user=" + config.Env.DbUser
	connString += " password=" + config.Env.DbPassword
	connString += " dbname=" + config.Env.DbName
	connString += " port=" + config.Env.DbPort
	connString += " sslmode=" + config.Env.DbSslMode

	return connString
}
