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

func updateUser(c *gin.Context) {
	id := c.Param("id")
	user := map[string]interface{}{}
	c.BindJSON(&user)

	updatedUser, err := services.UpdateUser(id, user)

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "fail",
			"message": "User not found or some error happened",
			"detail":  err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"user": updatedUser,
	})
}

// SetupUserRoute setup user route
func setupUserRoute(c *gin.Engine) {
	userRoute := c.Group("/api/users")
	{
		userRoute.GET("/", getUserInfo)
		userRoute.POST("/", createUser)
		userRoute.PUT("/:id", updateUser)
	}
}
