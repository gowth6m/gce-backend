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
	routes.IntroRoutes(router)
	versionControlled := router.Group("/" + configs.AppConfig().App.ApiVersion)
	{
		routes.DefaultRoutes(versionControlled)
		routes.UserRoutes(versionControlled)
	}
	router.Run(configs.AppConfig().App.Host + configs.AppConfig().App.Port)
	database.DisconnectFromMongoDB()
}
