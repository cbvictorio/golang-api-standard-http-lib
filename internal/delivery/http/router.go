package http

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, userHandler *UserHandler) {

	// Public routes
	r.POST("/sign-up", userHandler.SignUp)

	// Protected routes (we'll add middleware later)
	r.GET("/monitor", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "alive"})
	})
}
