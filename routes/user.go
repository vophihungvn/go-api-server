package routes

import (
	services "services"

	"github.com/gin-gonic/gin"
)

func getUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": services.GetUserList(),
	})
}

// SetupUserRoute setup user route
func setupUserRoute(c *gin.Engine) {
	userRoute := c.Group("/api/users")
	{
		userRoute.GET("/", getUserInfo)
	}
}
