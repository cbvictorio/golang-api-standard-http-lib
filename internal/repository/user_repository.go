package repository

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"golang-api-standard-http-lib/internal/domain"
	postgresRepository "golang-api-standard-http-lib/internal/repository/postgres"
)

type UserRepositoryAbstraction interface {
	GetByEmail(email string) (*domain.User, error)
	Create(user domain.User) error
}

type UserRepository struct {
	postgresClient *postgresRepository.PostgresClient
}

func NewUserRepository(pgClient *postgresRepository.PostgresClient) *UserRepository {
	return &UserRepository{postgresClient: pgClient}
}

func (userRepository *UserRepository) GetByEmail(email string) (*domain.User, error) {
	user := &domain.User{Email: email}

	result := userRepository.postgresClient.Where("email = ?", email).First(&user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		slog.Error("Something went wrong while retrieving the user", "error", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (userRepository *UserRepository) Create(userInput *domain.User) error {
	result := userRepository.postgresClient.Create(&userInput)
	return result.Error
}
