package routes

import (
	"github.com/gin-gonic/gin"
)

func DefaultRoutes(group *gin.RouterGroup) {

	group.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the GreatComCatEngineering API",
			"version": "1.0.0",
			"author":  "GreatComCatEngineering",
		})
	})

}
