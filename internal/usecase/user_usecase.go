package usecase

import (
	"errors"
	"golang-api-standard-http-lib/internal/domain"
	repository "golang-api-standard-http-lib/internal/repository/postgres"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) CreateUser(user domain.User) error {
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
