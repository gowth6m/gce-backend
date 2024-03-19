package routes

import (
	"github.com/gin-gonic/gin"
	"greatcomcatengineering.com/backend/services/user"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", func(c *gin.Context) {
			user.HandleCreateUser(c.Writer, c.Request)
		})

		userGroup.GET("/:id", func(c *gin.Context) {
			user.HandleGetUserByEmail(c.Writer, c.Request)
		})

		userGroup.GET("/all", func(c *gin.Context) {
			user.HandleGetAllUsers(c.Writer, c.Request)
		})
	}
}
