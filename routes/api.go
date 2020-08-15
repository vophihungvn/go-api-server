package routes

import "github.com/gin-gonic/gin"

// SetupRoute setup API routes
func SetupRoute(r *gin.Engine) {
	setupUserRoute(r)
	setupProductRoute(r)
}
