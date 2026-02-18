package http

import (
	"net/http"

	"log/slog"

	"github.com/gin-gonic/gin"

	"golang-api-standard-http-lib/internal/usecase"
	"golang-api-standard-http-lib/pkg"
)

func (h *UserHandler) Login(context *gin.Context) {
	var body usecase.UserCredentialsDTO
	context.Bind(&body)

	user, err := h.userService.AuthenticateUser(body.Email, body.Password)

	// validate a possible server error
	if err != nil {
		msg := err.Error()
		context.AbortWithStatusJSON(http.StatusBadRequest, pkg.HttpAppErrorResponse(msg))
		return
	}

	// if user does not exist
	if user == nil {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			pkg.HttpAppErrorResponse(pkg.ErrorMessage(err, "Invalid credentials")),
		)

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})

}

func (h *UserHandler) SignUp(context *gin.Context) {
	var body usecase.NewUserDTO
	context.Bind(&body)

	userExists, err := h.userService.DoesEmailExist(body.Email)

	// error happened
	if err != nil {
		msg := "Something went wrong while creating the user"
		slog.Error(msg, "error", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, pkg.HttpAppErrorResponse(msg))
		return
	}

	// user already exists
	if userExists {
		context.JSON(
			http.StatusBadRequest,
			pkg.HttpAppErrorResponse("This email address is already associated with an account."),
		)

		return
	}

	userCreationError := h.userService.CreateUser(&body)

	if userCreationError != nil {
		msg := "Something went wrong while creating the user"
		slog.Error(msg, "error", userCreationError.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, pkg.HttpAppErrorResponse(msg))
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}
