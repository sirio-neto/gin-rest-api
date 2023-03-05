package main

import (
	"github.com/sirio-neto/gin-rest-api/database"
	config "github.com/sirio-neto/gin-rest-api/environment"
	"github.com/sirio-neto/gin-rest-api/routes"
)

func main() {
	config.InitEnvironmentConfig()
	database.ConnDB()

	routes.HandleRequests()
}
