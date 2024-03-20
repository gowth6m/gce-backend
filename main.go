package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/routes"
	"log"
)

// @title Great Comcat Engineering API
// @version 1
// @description This is the API for the Great Comcat Engineering project.
// @host localhost:8080
// @BasePath /v0
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Connect to MongoDB
	database.ConnectToMongoDB()
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"greatcomcatengineering.com", "localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Routes
	routes.IntroRoutes(router)
	routes.SwaggerRoutes(router)
	versionControlled := router.Group("/" + configs.AppConfig().App.ApiVersion)
	{
		routes.DefaultRoutes(versionControlled)
		routes.UserRoutes(versionControlled)
		routes.ProductRoutes(versionControlled)
	}

	router.Run(configs.AppConfig().App.Host + configs.AppConfig().App.Port)
	database.DisconnectFromMongoDB()
}
