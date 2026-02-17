package http

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, userHandler *UserHandler) {

	r.POST("/login", userHandler.Login)
	r.POST("/sign-up", userHandler.SignUp)

}
