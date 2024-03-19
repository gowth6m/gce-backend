package main

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/routes"
	"log"
)

func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	database.ConnectToMongoDB()
	router := gin.Default()
	v0 := router.Group("/v0")
	{
		routes.DefaultRoutes(v0)
		routes.UserRoutes(v0)
	}
	router.Run("localhost:8080")
	database.DisconnectFromMongoDB()
}
