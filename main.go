package main

import (
	"midterm/database"
	"midterm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDB()

	r := gin.Default()

	routes.OrderRoutes(r, db)

	r.Run(":8080")
}
