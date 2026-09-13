package main

import (
	"mini-project-pbi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectDatabase()
	_ = db

	router := gin.Default()

	router.Run(":8080")
}
