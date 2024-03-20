package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/routes"
	"log"
	"net/http"
)

// @title Great Comcat Engineering API
// @version 1
// @description This is the API for the Great Comcat Engineering project.
// @host gce-backend.vercel.app
// @BasePath /v0
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Handler(w http.ResponseWriter, r *http.Request) {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Connect to MongoDB
	database.ConnectToMongoDB()
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"greatcomcatengineering.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
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
	router.ServeHTTP(w, r)
}
