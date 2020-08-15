package main

import (
	routes "routes"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	routes.SetupRoute(r)
	r.Run()
}
