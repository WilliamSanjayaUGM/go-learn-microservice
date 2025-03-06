package main

import (
	"go-learn-api/controllers"
	"go-learn-api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	models.ConnectDatabase()

	router.POST("/login", controllers.LoginUser)

	log.Println("Server started!")
	router.Run("localhost:9090")
}
