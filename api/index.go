package handler

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/services/user"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	router := gin.Default()

	router.GET("/user/:id", func(c *gin.Context) {
		user.HandleGetUserByEmail(c.Writer, c.Request)
	})

	router.POST("/user/create", func(c *gin.Context) {
		user.HandleCreateUser(c.Writer, c.Request)
	})

	// Use Gin to handle the request
	router.ServeHTTP(w, r)
}
