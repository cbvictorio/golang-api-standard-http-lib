package http

import (
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"

	"golang-api-standard-http-lib/internal/usecase"
	"golang-api-standard-http-lib/pkg"
)

type UserHandler struct {
	userService *usecase.UserService
}

func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(context *gin.Context) {
	var body usecase.UserCredentialsDTO
	context.Bind(&body)

	user, err := h.userService.GetByEmail(body.Email)

	// validate a possible server error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// if user does not exist
	if user == nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, pkg.AppError("Invalid username or password"))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"name":      user.Name,
		"email":     user.Email,
		"role":      user.Role,
		"createdAt": user.CreatedAt,
	})

}

func (h *UserHandler) SignUp(context *gin.Context) {
	var body usecase.NewUserDTO

	context.Bind(&body)

	user, err := h.userService.GetByEmail(body.Email)

	// error happened
	if err != nil {
		msg := "Something went wrong while creating the user"
		slog.Error(msg, "error", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, pkg.AppError(msg))
		return
	}

	// user already exists
	if user != nil {
		context.JSON(
			http.StatusBadRequest,
			pkg.AppError("This email address is already associated with an account."),
		)

		return
	}

	userCreationError := h.userService.CreateUser(&body)

	if userCreationError != nil {
		msg := "Something went wrong while creating the user"
		slog.Error(msg, "error", userCreationError.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, pkg.AppError(msg))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

}
