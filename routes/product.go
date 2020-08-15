package routes

import "github.com/gin-gonic/gin"

func handleGetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"products": []string{"product 1"},
	})
}

// SetupProductRoute handle products route
func setupProductRoute(c *gin.Engine) {
	productRoute := c.Group("/api/products")
	{
		productRoute.GET("/", handleGetProducts)
	}
}
