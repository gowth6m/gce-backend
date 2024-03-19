package handler

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/routes"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
	router.ServeHTTP(w, r)
}
