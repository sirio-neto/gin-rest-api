package main

import (
	"github.com/sirio-neto/gin-rest-api/database"
	"github.com/sirio-neto/gin-rest-api/routes"
)

func main() {
	database.ConnDB()

	routes.HandleRequests()
}
