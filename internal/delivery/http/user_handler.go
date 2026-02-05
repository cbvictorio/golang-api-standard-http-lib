package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"golang-api-standard-http-lib/internal/domain"
	"golang-api-standard-http-lib/internal/usecase"
	"golang-api-standard-http-lib/pkg"
)

type UserHandler struct {
	userService *usecase.UserService
}

func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) SignUp(context *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}
	context.Bind(&body)

	password, err := pkg.GenerateHashFromPassword(body.Password)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := domain.User{
		ID:        uuid.NewString(),
		Name:      body.Name,
		Email:     body.Email,
		Role:      domain.RoleCustomer,
		CreatedAt: time.Now(),
		Password:  password,
	}

	err = h.userService.Create(user)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
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
			"message": "user with that email was not found (BACKEND)",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user found",
	})

}
