package main

import (
	"os"

	"mini-project-pbi/database"
	"mini-project-pbi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDatabase()

	router := gin.Default()

	routes.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
