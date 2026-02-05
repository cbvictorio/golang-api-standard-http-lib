package usecase

import (
	"errors"
	"golang-api-standard-http-lib/internal/domain"
	repository "golang-api-standard-http-lib/internal/repository/postgres"
	"log/slog"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) Create(user domain.User) error {
	result := repository.PostgresClient.Create(&user)

	if result.Error != nil {
		errorMessage := "There was an error, please try again later"

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			errorMessage = "That username/email is already registered, please try a different one"
		}

		return errors.New(errorMessage)
	}

	return result.Error
}

func (userService *UserService) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{Email: email}
	result := repository.PostgresClient.Where("email = ?", email).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		slog.Error("Something went wrong while retrieving the user by email", "error", result.Error)
		return nil, result.Error
	}

	return user, nil
}
