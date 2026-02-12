package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-api-standard-http-lib/internal/usecase"
)

type UserHandler struct {
	userService *usecase.UserService
}

func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(context *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	context.Bind(&body)

	user, err := h.userService.GetByEmail(body.Email)

	// validate a possible server error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// validate existing user
	if user == nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "user with that email was not found",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user found",
	})

}
