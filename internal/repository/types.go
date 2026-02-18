package repository

import (
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
