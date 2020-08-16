package routes

import (
	"models"
	services "services"

	"github.com/gin-gonic/gin"
)

func getUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": services.GetUserList(),
	})
}

func createUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)

	services.CreateUser(user)

	c.JSON(200, gin.H{
		"status": "success",
	})
}

// SetupUserRoute setup user route
func setupUserRoute(c *gin.Engine) {
	userRoute := c.Group("/api/users")
	{
		userRoute.GET("/", getUserInfo)
		userRoute.POST("/", createUser)
	}
}
